package web

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (s *Server) uploadsHead(c echo.Context) error {
	return s.assetsHeadImpl(s.config.UploadsRoot, c)
}

func (s *Server) assetsHead(c echo.Context) error {
	return s.assetsHeadImpl(s.config.AssetsRoot, c)
}

func (s *Server) assetsHeadImpl(p string, c echo.Context) error {
	name := filepath.ToSlash(filepath.Clean(strings.TrimPrefix(c.Param("*"), "/")))

	uploadsRoot, err := filepath.Abs(p)
	if err != nil {
		return err
	}

	uploadPath, err := filepath.Abs(path.Join(uploadsRoot, name))
	if err != nil {
		return err
	}

	if !startsWith(uploadPath, uploadsRoot) {
		return errors.New("not in directory")
	}

	stat, err := os.Stat(uploadPath)
	if err != nil {
		return err
	}

	c.Response().Header().Add("Content-Length", fmt.Sprintf("%d", stat.Size()))

	return c.String(http.StatusOK, "")
}
