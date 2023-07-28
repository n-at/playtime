package web

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
)

func (s *Server) contextCustomizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var session storage.Session

		cookie, err := c.Cookie(SessionCookieName)
		if err == nil {
			session, err = s.storage.SessionGetById(cookie.Value)
			if err != nil {
				log.Warnf("unable to get session by sessionId: %s", err)
			}
		}

		ensembleContext := &PlaytimeContext{
			Context: c,
			session: &session,
		}

		return next(ensembleContext)
	}
}

///////////////////////////////////////////////////////////////////////////////

func (s *Server) authenticationRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		context := c.(*PlaytimeContext)

		if context.session == nil || len(context.session.UserId) == 0 {
			return c.Redirect(http.StatusFound, "/login")
		}

		user, err := s.storage.UserFindById(context.session.UserId)
		if err != nil {
			log.Warnf("unable to get user by session userId: %s", err)
			return c.Redirect(http.StatusFound, "/login")
		}
		if len(user.Id) == 0 {
			log.Warn("user by session not found")
			return c.Redirect(http.StatusFound, "/login")
		}

		context.user = &user

		return next(context)
	}
}

func (s *Server) settingsRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		context := c.(*PlaytimeContext)

		settings, err := s.storage.SettingsGetByUserId(context.user.Id)
		if err != nil {
			return err
		}

		context.settings = &settings

		return next(context)
	}
}
