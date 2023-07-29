package web

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
)

func (s *Server) users(c echo.Context) error {
	users, err := s.storage.UserFindAll()
	if err != nil {
		log.Errorf("users get all error: %s", err)
		return err
	}

	context := c.(*PlaytimeContext)

	return c.Render(http.StatusOK, "users", pongo2.Context{
		"user":  context.user,
		"users": users,
		"done":  c.Param("done"),
	})
}

func (s *Server) userNewForm(c echo.Context) error {
	context := c.(*PlaytimeContext)
	return c.Render(http.StatusOK, "user_new", pongo2.Context{
		"user": context.user,
	})
}

func (s *Server) userNewSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	log.Infof("userNewSubmit")

	password, err := storage.EncryptPassword(c.FormValue("password"))
	if err != nil {
		log.Errorf("userNewSubmit password encryption error: %s", err)
		return err
	}

	user := storage.User{
		Login:    c.FormValue("login"),
		Password: password,
		Active:   c.FormValue("active") == "1",
		Admin:    c.FormValue("admin") == "1",
	}
	if _, err := s.storage.UserSave(user); err != nil {
		log.Errorf("userNewSubmit user save error: %s", err)
		return c.Render(http.StatusOK, "user_new", pongo2.Context{
			"user":         context.user,
			"user_control": user,
			"error":        err,
		})
	}

	return c.Redirect(http.StatusFound, "/users?done=1")
}

func (s *Server) userEditForm(c echo.Context) error {
	context := c.(*PlaytimeContext)
	return c.Render(http.StatusOK, "user_edit", pongo2.Context{
		"user":         context.user,
		"user_control": context.userControl,
	})
}

func (s *Server) userEditSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	log.Infof("userEditSubmit %s", context.userControl.Id)

	user := context.userControl
	user.Login = c.FormValue("login")
	user.Active = c.FormValue("active") == "1"
	user.Admin = c.FormValue("admin") == "1"

	password := c.FormValue("password")
	if len(password) != 0 {
		password, err := storage.EncryptPassword(password)
		if err != nil {
			log.Errorf("userEditSubmit user %s password encryption error: %s", context.userControl.Id, err)
			return err
		}
		user.Password = password
	}

	if _, err := s.storage.UserSave(*user); err != nil {
		log.Errorf("userEditSubmit user %s save error: %s", context.userControl.Id, err)
		return c.Render(http.StatusOK, "user_edit", pongo2.Context{
			"user":         context.user,
			"user_control": user,
			"error":        err,
		})
	}

	if !user.Active {
		if err := s.storage.SessionDeleteByUserId(user.Id); err != nil {
			log.Warnf("userEditSubmit unable to delete inactive user sessions: %s", err)
		}
	}

	return c.Redirect(http.StatusFound, "/users?done=1")
}

func (s *Server) userDeleteForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	if context.user.Id == context.userControl.Id {
		return errors.New("cannot delete self")
	}

	return c.Render(http.StatusOK, "user_delete", pongo2.Context{
		"user":         context.user,
		"user_control": context.userControl,
	})
}

func (s *Server) userDeleteSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	log.Infof("userDeleteSubmit %s", context.userControl.Id)

	if context.user.Id == context.userControl.Id {
		return errors.New("cannot delete self")
	}
	if err := s.storage.UserDeleteById(context.userControl.Id); err != nil {
		log.Errorf("userDeleteSubmit user %s delete error: %s", context.userControl.Id, err)
		return err
	}

	if err := s.storage.SessionDeleteByUserId(context.userControl.Id); err != nil {
		log.Warnf("userDeleteSubmit unable to delete deleted user sessions: %s", err)
	}

	return c.Redirect(http.StatusFound, "/users?done=1")
}
