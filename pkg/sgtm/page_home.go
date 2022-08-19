package sgtm

import (
	"net/http"
	"time"

	"github.com/gobuffalo/packr/v2"
	"go.uber.org/zap"
	"moul.io/sgtm/pkg/sgtmpb"
)

func (svc *Service) homePage(box *packr.Box) func(w http.ResponseWriter, r *http.Request) {
	tmpl := loadTemplates(box, "base.tmpl.html", "home.tmpl.html")
	return func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		data, err := svc.newTemplateData(w, r)
		if err != nil {
			svc.errRenderHTML(w, r, err, http.StatusUnprocessableEntity)
			return
		}
		// custom
		data.PageKind = "home"

		// tracking
		{
			viewEvent := sgtmpb.Post{AuthorID: data.UserID, Kind: sgtmpb.Post_ViewHomeKind}
			err = svc.store.CreatePost(&viewEvent)
			if err != nil {
				data.Error = "Cannot write activity: " + err.Error()
			} else {
				svc.logger.Debug("new view home", zap.Any("event", &viewEvent))
			}
		}

		// last tracks
		{
			limit := 50
			if data.UserID == 0 {
				limit = 10
			}
			data.Home.LastTracks, err = svc.store.GetPostList(limit)
			if err != nil {
				dat