package web

import (
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
)

func (s *Server) index(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/games")
}

func (s *Server) loginForm(c echo.Context) error {
	context := c.(*PlaytimeContext)
	if len(context.session.UserId) != 0 {
		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "login", pongo2.Context{})
}

func (s *Server) loginSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)
	if context.user != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	login := c.FormValue("login")
	password := c.FormValue("password")

	log.Infof("loginSubmit: %s", login)

	user, err := s.storage.UserFindByLogin(login)
	if err != nil || len(user.Id) == 0 {
		log.Errorf("loginSubmit user %s get user error: %s", login, err)
		return c.Render(http.StatusOK, "login", pongo2.Context{
			"login": login,
			"error": "Incorrect login or password",
		})
	}

	if !storage.CheckPassword(password, user.Password) {
		log.Warnf("loginSubmit user %s check password error", login)
		return c.Render(http.StatusOK, "login", pongo2.Context{
			"login": login,
			"error": "Incorrect login or password",
		})
	}

	session := storage.Session{
		UserId: user.Id,
	}

	session, err = s.storage.SessionSave(session)
	if err != nil {
		log.Errorf("loginSubmit user %s session creation error: %s", login, err)
		return c.Render(http.StatusOK, "login", pongo2.Context{
			"login": login,
			"error": err.Error(),
		})
	}

	context.SetSessionId(session.Id)

	return c.Redirect(http.StatusFound, "/")
}

func (s *Server) logout(c echo.Context) error {
	context := c.(*PlaytimeContext)
	context.DeleteSessionId()
	return c.Redirect(http.StatusFound, "/login")
}
