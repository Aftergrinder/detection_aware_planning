
package sgtm

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/dustin/go-humanize"
	"github.com/gobuffalo/packr/v2"
	striptags "github.com/grokify/html-strip-tags-go"
	"github.com/hako/durafmt"
	"github.com/kyokomi/emoji/v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"go.uber.org/zap"

	"moul.io/sgtm/internal/sgtmversion"
	"moul.io/sgtm/pkg/sgtmpb"
)

func (svc *Service) newTemplateData(w http.ResponseWriter, r *http.Request) (*templateData, error) {
	data := templateData{
		Title:            "SGTM",
		Date:             time.Now(),
		Opts:             svc.opts.Filtered(),
		Lang:             "en", // FIXME: dynamic
		Request:          r,
		Service:          svc,
		PageKind:         "other",
		ReleaseVersion:   sgtmversion.Version,
		ReleaseVcsRef:    sgtmversion.VcsRef,
		ReleaseBuildDate: sgtmversion.BuildDate,
	}
	if svc.opts.DevMode {
		data.Title += " (dev)"
	}

	if cookie, err := r.Cookie(oauthTokenCookie); err == nil {
		data.JWTToken = cookie.Value
		var err error
		data.Claims, err = svc.parseJWTToken(data.JWTToken)
		if err != nil {
			return nil, fmt.Errorf("parse jwt token: %w", err)
		}
		var user *sgtmpb.User
		user, err = svc.store.GetUserRecentPost(data.Claims.Session.UserID)
		if err != nil {
			svc.logger.Warn("load user from DB", zap.Error(err))
		}
		if user != nil {
			data.User = user
			data.UserID = user.ID
			data.IsAdmin = user.Role == "admin"
			// w.Header().Set("SGTM-User-ID", fmt.Sprintf("%d", user.ID))
			w.Header().Set("SGTM-User-Slug", user.Slug)
		}
	} else {
		w.Header().Set("SGTM-User-Slug", "-")
	}

	return &data, nil
}

func loadTemplates(box *packr.Box, filenames ...string) *template.Template {
	allInOne := ""
	templateName := ""
	for _, filename := range filenames {
		src, err := box.FindString("page_" + filename)
		if err != nil {
			panic(err)
		}
		allInOne += strings.TrimSpace(src) + "\n"
		templateName += filename
	}
	allInOne = strings.TrimSpace(allInOne)
	funcmap := sprig.FuncMap()
	funcmap["markdownify"] = func(input string) template.HTML {
		extensions := blackfriday.CommonExtensions | blackfriday.NoEmptyLineBeforeBlock
		unsafe := blackfriday.Run([]byte(input), blackfriday.WithExtensions(extensions))
		mdHTML := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		html := fmt.Sprintf(`<div class="markdownify">%s</div>`, string(mdHTML))
		return template.HTML(html)
	}
	funcmap["fromUnixNano"] = func(input int64) time.Time {
		return time.Unix(0, input)
	}
	funcmap["prettyURL"] = func(input string) string {
		u, err := url.Parse(input)
		if err != nil {
			return ""
		}
		if len(u.Path) > 10 {
			u.Path = u.Path[0:7] + "..."
		}
		shorten := fmt.Sprintf("%s%s", u.Host, u.Path)
		shorten = strings.TrimRight(shorten, "/")
		shorten = strings.TrimPrefix(shorten, "www.")
		return shorten
	}
	funcmap["emojify"] = func(input string) string {
		return emoji.Sprint(input)
	}
	funcmap["newline"] = func() string {
		return "\n"
	}