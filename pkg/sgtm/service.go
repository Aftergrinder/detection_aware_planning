package sgtm

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"moul.io/banner"

	"moul.io/sgtm/pkg/sgtmpb"
	"moul.io/sgtm/pkg/sgtmstore"
)

type Service struct {
	sgtmpb.UnimplementedWebAPIServer

	store         sgtmstore.Store
	logger        *zap.Logger
	opts          Opts
	ctx           context.Context
	cancel        func()
	StartedAt     time.Time
	errRenderHTML func(w http.ResponseWriter, r *http.Request, err error, status int)

	// drivers

	discord          discordDriver
	server           serverDriver
	processingWorker processingWorkerDriver
	ipfs             ipfsWrapper
	unittest         bool
}

// New constructor that initializes new Service
func New(store sgtmstore.Store, opts Opts) (*Service, error) {
	if err := opts.applyDefaults(); err != nil {
		return nil, err
	}
	fmt.Fpri