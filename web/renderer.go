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
)

var (
	TemplatesRoot      = "templates"
	TemplatesExtension = "twig"
	TemplatesDebug     = false
)

type pongo2Renderer struct{}

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
	if TemplatesDebug {
		t, err = pongo2.FromFile(resolveTemplateName(name))
	} else {
		t, err = pongo2.FromCache(resolveTemplateName(name))
	}
	if err != nil {
		return err
	}

	return t.ExecuteWriter(ctx, w)
}

func resolveTemplateName(n string) string {
	return fmt.Sprintf("%s%c%s.%s", TemplatesRoot, os.PathSeparator, n, TemplatesExtension)
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
