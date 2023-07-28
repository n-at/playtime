package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"playtime/storage"
	"time"
)

type PlaytimeContext struct {
	echo.Context
	session     *storage.Session
	user        *storage.User
	game        *storage.Game
	saveState   *storage.SaveState
	loadBatch   *storage.LoadBatch
	settings    *storage.Settings
	userControl *storage.User
}

func (c *PlaytimeContext) GetSessionId() string {
	cookie, err := c.Cookie(SessionCookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (c *PlaytimeContext) SetSessionId(id string) {
	cookie := &http.Cookie{
		Name:    SessionCookieName,
		Value:   id,
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)
}

func (c *PlaytimeContext) DeleteSessionId() {
	cookie := &http.Cookie{
		Name:    SessionCookieName,
		Value:   "",
		Expires: time.Now().Add(-24 * time.Hour),
	}
	c.SetCookie(cookie)
}
