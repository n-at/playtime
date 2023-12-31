package web

import (
	"errors"
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

func (s *Server) userControlAccessRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		context := c.(*PlaytimeContext)

		if context.user == nil {
			return errors.New("authentication required")
		}
		if !context.user.CanControlUsers() {
			return errors.New("user control access denied")
		}

		return next(c)
	}
}

func (s *Server) userControlRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("user_id")
		if len(userId) == 0 {
			return errors.New("user id required")
		}

		user, err := s.storage.UserFindById(userId)
		if err != nil {
			return err
		}
		if len(user.Id) == 0 {
			return errors.New("user not found")
		}

		context := c.(*PlaytimeContext)
		context.userControl = &user

		return next(context)
	}
}

func (s *Server) gameRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		gameId := c.Param("game_id")
		if len(gameId) == 0 {
			return errors.New("game id required")
		}

		game, err := s.storage.GameGetById(gameId)
		if err != nil {
			return err
		}
		if len(game.Id) == 0 {
			return errors.New("game not found")
		}

		context := c.(*PlaytimeContext)
		if context.user == nil {
			return errors.New("authentication required")
		}
		if game.UserId != context.user.Id {
			return errors.New("game belongs to different user")
		}

		context.game = &game

		return next(context)
	}
}

func (s *Server) netplayGameRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		gameId := c.Param("game_id")
		if len(gameId) == 0 {
			return errors.New("game id required")
		}

		netplaySessionId := c.Param("netplay_session_id")
		if len(netplaySessionId) == 0 {
			return errors.New("netplay session id required")
		}

		game, err := s.storage.GameGetById(gameId)
		if err != nil {
			return err
		}
		if len(game.Id) == 0 {
			return errors.New("game not found")
		}
		if game.NetplaySessionId != netplaySessionId {
			return errors.New("netplay session id mismatch")
		}

		context := c.(*PlaytimeContext)
		context.game = &game

		return next(context)
	}
}

func (s *Server) uploadBatchRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uploadBatchId := c.Param("upload_batch_id")
		if len(uploadBatchId) == 0 {
			return errors.New("upload batch id required")
		}

		uploadBatch, err := s.storage.UploadBatchGetById(uploadBatchId)
		if err != nil {
			return err
		}
		if len(uploadBatch.Id) == 0 {
			return errors.New("upload batch not found")
		}

		context := c.(*PlaytimeContext)
		if context.user == nil {
			return errors.New("authentication required")
		}
		if context.user.Id != uploadBatch.UserId {
			return errors.New("upload batch belongs to different user")
		}

		context.uploadBatch = &uploadBatch

		return next(context)
	}
}

func (s *Server) saveStateRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		saveStateId := c.Param("save_state_id")
		if len(saveStateId) == 0 {
			return errors.New("save state id required")
		}

		saveState, err := s.storage.SaveStateGetById(saveStateId)
		if err != nil {
			return err
		}
		if len(saveState.Id) == 0 {
			return errors.New("save state not found")
		}

		context := c.(*PlaytimeContext)
		if context.user == nil {
			return errors.New("authentication required")
		}
		if context.user.Id != saveState.UserId {
			return errors.New("save state belongs to different user")
		}

		game, err := s.storage.GameGetById(saveState.GameId)
		if err != nil {
			return err
		}
		if len(game.Id) == 0 {
			return errors.New("save state game not found")
		}
		if context.user.Id != game.UserId {
			return errors.New("save state game belongs to different user")
		}

		context.game = &game
		context.saveState = &saveState

		return next(context)
	}
}
