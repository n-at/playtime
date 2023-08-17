package web

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/web/gamesession"
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
		if session.CountClients() == 0 {
			s.gameSessions.DeleteSession(sessionId)
			continue
		}

		for _, clientId := range session.GetClients() {
			if !session.IsHeartbeatReceived(clientId) {
				if err := session.DisconnectAndRemove(clientId); err != nil {
					log.Warnf("unable to disconnect client %s from session %s: %s", clientId, sessionId, err)
					session.RemoveClient(clientId)
				}
				continue
			}

			s.heartbeatPool.Add(func() {
				msg := gamesession.MessageOutgoing{
					Type: gamesession.MessageTypeHeartbeat,
					Heartbeat: &gamesession.MessageOutgoingHeartbeat{
						ClientId: clientId,
					},
				}
				session.SetHeartbeatReceived(clientId, false)
				session.Send(clientId, msg)
			})
		}
	}
}