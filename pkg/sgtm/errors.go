package sgtm

import (
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

func (svc *Service) errRender(w http.ResponseWriter, r *http.Request, err error, status int) {
	renderer :