package sgtm

import (
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

func (svc *Service) errRender(w http.ResponseWriter, r *http.Request, err error, status int) {
	renderer := errToResponse(err, status)
	svc.logger.
		WithOptions(zap.AddCallerSkip(1)).
		Warn(
			"user error",
			zap.String("title", renderer.Title),
			zap.Error(err),
		)
	if err := render.Render(w, r, renderer); err != nil {
		svc.logger.Warn("cannot render error", zap.Error(err))
	}
}

func errToResponse(err error,