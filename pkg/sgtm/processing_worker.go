
package sgtm

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"moul.io/banner"
	"moul.io/sgtm/pkg/sgtmpb"
)

type processingWorkerDriver struct {
	started bool
	wg      *sync.WaitGroup

	trackMigrations []func(*sgtmpb.Post, *gorm.DB) error
}

func (svc *Service) StartProcessingWorker() error {
	// init
	{
		fmt.Fprintln(os.Stderr, banner.Inline("processing-worker"))
		svc.logger.Debug("starting processing-worker")
		svc.setupMigrations()
		svc.processingWorker.wg = &sync.WaitGroup{}
		svc.processingWorker.wg.Add(1)
		defer svc.processingWorker.wg.Done()
		svc.processingWorker.started = true
	}

	// loop
	for i := 0; ; i++ {
		if err := svc.processingLoop(i); err != nil {
			return err
		}

		select {
		// FIXME: add a channel to get instant worker task
		case <-time.After(30 * time.Second):
		case <-svc.ctx.Done():
			return nil
		}
	}
}

func (svc *Service) CloseProcessingWorker(err error) {
	svc.logger.Debug("closing processingWorker", zap.Bool("was-started", svc.processingWorker.started), zap.Error(err))
	svc.cancel()
	if svc.processingWorker.started {
		svc.processingWorker.wg.Wait()
		svc.logger.Debug("processing-worker closed")
	}
}

func (svc *Service) processingLoop(i int) error {
	before := time.Now()

	// track migrations
	{
		var outdated []*sgtmpb.Post
		err := svc.store.DB().
			Where(sgtmpb.Post{Kind: sgtmpb.Post_TrackKind}).
			Where("processing_error IS NULL OR processing_error == ''").
			Where("processing_version IS NULL OR processing_version < ?", len(svc.processingWorker.trackMigrations)).
			Preload("Author").
			Find(&outdated).
			Error
		if err != nil {
			return fmt.Errorf("failed to fetch tracks that need to be processed: %w", err)
		}

		err = svc.store.DB().Omit(clause.Associations).Transaction(func(tx *gorm.DB) error {
			for _, entryPtr := range outdated {
				entry := entryPtr
				version := 1
				for _, migration := range svc.processingWorker.trackMigrations {