package web

import (
	"errors"
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"playtime/web/localization"
)

type pongo2Renderer struct {
	config *Configuration
}

func newPongo2Renderer(config *Configuration) pongo2Renderer {
	return pongo2Renderer{config: config}
}

func (r pongo2Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var ctx pongo2.Context
	var ok bool
	if data != nil {
		ctx, ok = data.(pongo2.Context)
		if !ok {
			return errors.New("no pongo2.Context data was passed")
		}
	}

	var t *pongo2.Template
	var err error
	if r.config.TemplatesDebug {
		t, err = pongo2.FromFile(r.resolveTemplateName(name))
	} else {
		t, err = pongo2.FromCache(r.resolveTemplateName(name))
	}
	if err != nil {
		return err
	}

	//l10n
	lang := localization.DefaultLanguageCode
	langCookie, err := c.Cookie("playtime-l10n")
	if err == nil && localization.Exists(langCookie.Value) {
		lang = langCookie.Value
	}
	ctx["localization_list"] = localization.List()
	ctx["localization_lang"] = lang
	ctx["loc"] = func(s string, args ...any) string {
		return localization.Localize(lang, s, args)
	}

	return t.ExecuteWriter(ctx, w)
}

func (r pongo2Renderer) resolveTemplateName(n string) string {
	return fmt.Sprintf("%s%c%s.%s", r.config.TemplatesRoot, os.PathSeparator, n, r.config.TemplatesExtension)
}

func httpErrorHandler(e error, c echo.Context) {
	code := http.StatusInternalServerError
	if httpError, ok := e.(*echo.HTTPError); ok {
		code = httpError.Code
	}

	log.WithFields(log.Fields{
		"method": c.Request().Method,
		"uri":    c.Request().URL,
		"error":  e,
	}).Error("request error")

	if err := c.Render(code, "error", pongo2.Context{"error": e}); err != nil {
		log.Errorf("error page render error: %s", err)
	}
}
