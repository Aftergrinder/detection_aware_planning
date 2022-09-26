
package sgtm

import (
	"net/http"
	"time"

	"github.com/gobuffalo/packr/v2"
)

func (svc *Service) rssPage(box *packr.Box) func(w http.ResponseWriter, r *http.Request) {
	tmpl := loadTemplates(box, "rss.tmpl.xml")
	return func(w http.ResponseWriter, r *http.Request) {