package web

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"net/http"
	"runtime/debug"
)

func (s *Server) version(c echo.Context) error {
	context := c.(*PlaytimeContext)

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return errors.New("unable to read build info")
	}

	return c.Render(http.StatusOK, "version", pongo2.Context{
		"_csrf_token": c.Get("csrf"),
		"user":        context.user,
		"build":       buildInfo,
	})
}
