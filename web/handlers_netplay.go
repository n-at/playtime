package web

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
)

func (s *Server) netplay(c echo.Context) error {
	context := c.(*PlaytimeContext)

	game := context.game

	if !s.config.NetplayEnabled {
		return errors.New("netplay not available")
	}
	if !game.NetplayEnabled || len(game.NetplaySessionId) == 0 {
		return errors.New("netplay for game not available")
	}

	return c.Render(http.StatusOK, "netplay", pongo2.Context{
		"game":                  game,
		"controls":              s.findNetplayControls(context),
		"netplay_turn_url":      s.config.TurnServerUrl,
		"netplay_turn_user":     s.config.TurnServerUser,
		"netplay_turn_password": s.config.TurnServerPassword,
	})
}

func (s *Server) netplayWS(c echo.Context) error {
	return nil //TODO
}

func (s *Server) netplayHeartbeat() {
	log.Debug("sending netplay heartbeats")

	for _, sessionId := range s.gameSessions.GetSessions() {
		session := s.gameSessions.GetSession(sessionId)
		if session == nil {
			continue
		}
		if session.CountPlayers() == 0 {
			s.gameSessions.DeleteSession(sessionId)
			continue
		}

		for _, playerId := range session.GetPlayers() {
			if !session.IsHeartbeatReceived(playerId) {
				if err := session.DisconnectAndRemove(playerId); err != nil {
					log.Warnf("unable to disconnect player %s from session %s: %s", playerId, sessionId, err)
					session.RemovePlayer(playerId)
				}
				continue
			}

			s.heartbeatPool.Add(func() {
				session.SetHeartbeatReceived(playerId, false)
				session.Send(playerId, nil) //TODO send actual heartbeat message
			})
		}
	}
}

///////////////////////////////////////////////////////////////////////////////

func (s *Server) findNetplayControls(context *PlaytimeContext) storage.EmulatorControls {
	game := context.game

	if game == nil || len(game.Platform) == 0 {
		return storage.EmulatorControls{}
	}

	if context.session != nil && len(context.session.UserId) != 0 {
		userSettings, err := s.storage.SettingsGetByUserId(context.session.UserId)
		if err != nil {
			log.Warnf("unable to get current user %s settings: %s", context.settings.UserId, err)
		} else {
			return userSettings.EmulatorSettings[game.Platform].Controls[0]
		}
	}

	userSettings, err := s.storage.SettingsGetByUserId(game.UserId)
	if err != nil {
		log.Warnf("unable to get user %s settings: %s", game.UserId, err)
	} else {
		return userSettings.EmulatorSettings[game.Platform].Controls[0]
	}

	return storage.DefaultEmulatorSettings(game.Platform).Controls[0]
}
