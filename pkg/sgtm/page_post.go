
package sgtm

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/gobuffalo/packr/v2"
	"go.uber.org/zap"
	"moul.io/godev"

	"moul.io/sgtm/pkg/sgtmpb"
)

func (svc *Service) postPage(box *packr.Box) func(w http.ResponseWriter, r *http.Request) {
	tmpl := loadTemplates(box, "base.tmpl.html", "post.tmpl.html")

	return func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		data, err := svc.newTemplateData(w, r)
		if err != nil {
			svc.errRenderHTML(w, r, err, http.StatusUnprocessableEntity)
			return
		}
		// custom
		data.PageKind = "post"
		postSlug := chi.URLParam(r, "post_slug")
		post, err := svc.store.GetPostBySlugOrID(postSlug)
		if err != nil {
			svc.error404Page(box)(w, r)
			return
		}
		data.Post.Post = post
		data.Post.Post.ApplyDefaults()

		if r.URL.Query().Get("format") == "json" {
			data.Post.Post.Filter()
			data.Post.Post.Author.Filter()
			fmt.Fprintln(w, godev.PrettyJSONPB(data.Post.Post))
			return
		}

		if r.Method == "POST" && data.UserID != 0 {
			validate := func() *sgtmpb.Post {
				if err := r.ParseForm(); err != nil {
					data.Error = err.Error()
					return nil
				}
				comment := sgtmpb.Post{
					Kind:         sgtmpb.Post_CommentKind,
					AuthorID:     data.UserID,
					Body:         strings.TrimSpace(r.Form.Get("comment")),
					Visibility:   sgtmpb.Visibility_Public,
					TargetPostID: data.Post.Post.ID,
				}
				if comment.Body == "" {
					return nil
				}
				return &comment
			}
			comment := validate()
			if comment != nil {
				err = svc.store.CreatePost(comment)
				if err != nil {
					svc.errRenderHTML(w, r, err, http.StatusUnprocessableEntity)
					return
				}
				svc.logger.Debug("comment created", zap.Any("post", comment))
				http.Redirect(w, r, data.Post.Post.CanonicalURL(), http.StatusFound)
				return
			}
		}

		// load comments
		{
			data.Post.Comments, err = svc.store.GetPostComments(data.Post.Post.ID)
			if err != nil {
				svc.errRenderHTML(w, r, err, http.StatusUnprocessableEntity)
				return
			}
		}

		// tracking
		{
			viewEvent := sgtmpb.Post{AuthorID: data.UserID, Kind: sgtmpb.Post_ViewPostKind, TargetPostID: data.Post.Post.ID}
			if err := svc.store.CreatePost(&viewEvent); err != nil {
				data.Error = "Cannot write activity: " + err.Error()
			} else {
				svc.logger.Debug("new view post", zap.Any("event", &viewEvent))
			}
		}

		// end of custom
		if svc.opts.DevMode {
			tmpl = loadTemplates(box, "base.tmpl.html", "post.tmpl.html")
		}
		data.Duration = time.Since(started)
		if err := tmpl.Execute(w, &data); err != nil {
			svc.errRenderHTML(w, r, err, http.StatusUnprocessableEntity)
			return
		}
	}
}

func (svc *Service) postMaintenancePage(box *packr.Box) func(w http.ResponseWriter, r *http.Request) {
	tmpl := loadTemplates(box, "base.tmpl.html", "dummy.tmpl.html")
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			shouldExtractBpm          = r.URL.Query().Get("extract_bpm") == "1"
			shouldDetectRelationships = r.URL.Query().Get("detect_relationships") == "1"
			shouldResyncSoundCloud    = r.URL.Query().Get("resync_soundcloud") == "1"
			shouldDL                  = shouldExtractBpm
			shouldDoSomething         = shouldExtractBpm || shouldDetectRelationships || shouldResyncSoundCloud
		)
		if !shouldDoSomething {
			svc.error404Page(box)(w, r)
			return
		}

		// common init
		started := time.Now()
		data, err := svc.newTemplateData(w, r)
		if err != nil {
			svc.errRenderHTML(w, r, err, http.StatusUnprocessableEntity)
			return
		}
		// custom
		if data.User == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		postSlug := chi.URLParam(r, "post_slug")
		post, err := svc.store.GetPostBySlugOrID(postSlug)
		if err != nil {
			svc.error404Page(box)(w, r)