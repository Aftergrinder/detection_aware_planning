package sgtm

import (
	"net/http"
	"time"

	packr "github.com/gobuffalo/packr/v2"
)

func (svc *Service) error404Page(box *packr.Box) func(w http.ResponseWriter, r *http.Request) {
	tmpl := loadTemplates(box, "base.tmpl.html", "err