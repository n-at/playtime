package web

import (
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
)

func (s *Server) settingsGeneralForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	return c.Render(http.StatusOK, "settings", pongo2.Context{
		"user":      context.user,
		"settings":  context.settings,
		"done":      c.QueryParam("done"),
		"languages": storage.Languages,
		"shaders":   storage.Shaders,
		"platforms": storage.Platforms,
	})
}

func (s *Server) settingsGeneralSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	log.Infof("generalSettingsSubmit for %s", context.user.Login)

	settings := context.settings
	settings.Language = c.FormValue("language")
	if _, err := s.storage.SettingsSave(*settings); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/settings?done=1")
}

func (s *Server) settingsByPlatformForm(c echo.Context) error {
	return nil
}

func (s *Server) settingsByPlatformSubmit(c echo.Context) error {
	return nil
}
