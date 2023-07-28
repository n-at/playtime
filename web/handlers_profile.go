package web

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
)

func (s *Server) profileForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	return c.Render(http.StatusOK, "profile", pongo2.Context{
		"user": context.user,
		"done": c.QueryParam("done"),
	})
}

func (s *Server) profileSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	log.Infof("profileSubmit for %s", context.user.Login)

	currentPassword := c.FormValue("password")
	if !storage.CheckPassword(currentPassword, context.user.Password) {
		log.Warnf("profileSubmit user %s check password error", context.user.Login)
		return errors.New("wrong current password")
	}

	password1 := c.FormValue("password1")
	password2 := c.FormValue("password2")
	if password1 != password2 {
		return errors.New("passwords does not match")
	}
	if len(password1) == 0 {
		return errors.New("password should not be empty")
	}

	passwordEncrypted, err := storage.EncryptPassword(password1)
	if err != nil {
		log.Errorf("profileSubmit user %s password encryption error: %s", context.user.Login, err)
		return err
	}

	user := context.user
	user.Password = passwordEncrypted
	if _, err := s.storage.UserSave(*user); err != nil {
		log.Errorf("profileSubmit user %s update error: %s", context.user.Login, err)
		return err
	}

	return c.Redirect(http.StatusFound, "/profile?done=1")
}
