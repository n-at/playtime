package web

import (
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
	"playtime/web/localization"
	"time"
)

func (s *Server) index(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/games")
}

func (s *Server) loginForm(c echo.Context) error {
	context := c.(*PlaytimeContext)
	if len(context.session.UserId) != 0 {
		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "login", pongo2.Context{
		"_csrf_token": c.Get("csrf"),
		"return":      c.FormValue("return"),
	})
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
			"_csrf_token": c.Get("csrf"),
			"login":       login,
			"error":       localization.LocalizeCtx(c, "login.incorrect"),
			"return":      c.FormValue("return"),
		})
	}

	if !storage.CheckPassword(password, user.Password) {
		log.Warnf("loginSubmit user %s check password error", login)
		return c.Render(http.StatusOK, "login", pongo2.Context{
			"_csrf_token": c.Get("csrf"),
			"login":       login,
			"error":       localization.LocalizeCtx(c, "login.incorrect"),
			"return":      c.FormValue("return"),
		})
	}

	session := storage.Session{
		Id:      storage.NewId(),
		UserId:  user.Id,
		Created: time.Now(),
	}

	if _, err = s.storage.SessionSave(session); err != nil {
		log.Errorf("loginSubmit user %s session creation error: %s", login, err)
		return c.Render(http.StatusOK, "login", pongo2.Context{
			"_csrf_token": c.Get("csrf"),
			"login":       login,
			"error":       err.Error(),
			"return":      c.FormValue("return"),
		})
	}

	context.SetSessionId(session.Id, c.IsTLS())

	ret := c.FormValue("return")
	if len(ret) == 0 {
		ret = "/"
	}

	return c.Redirect(http.StatusFound, ret)
}

func (s *Server) logout(c echo.Context) error {
	context := c.(*PlaytimeContext)

	if context.session != nil {
		if err := s.storage.SessionDeleteById(context.session.Id); err != nil {
			log.Errorf("logout session delete error: %s", err)
		}
	}

	context.DeleteSessionId()

	return c.Redirect(http.StatusFound, "/login")
}
