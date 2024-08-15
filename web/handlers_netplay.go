package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
	"playtime/web/gamesession"
)

func (s *Server) netplay(c echo.Context) error {
	pctx := c.(*PlaytimeContext)

	game := pctx.game

	if !s.config.NetplayEnabled {
		return errors.New("netplay not available")
	}
	if !game.NetplayEnabled || len(game.NetplaySessionId) == 0 {
		return errors.New("netplay for game not available")
	}

	user, err := s.findContextSessionUser(pctx)
	if err != nil {
		return err
	}

	if user != nil && user.Id == game.UserId {
		return c.Redirect(http.StatusFound, "/play/"+game.Id)
	}

	if user == nil && game.NetplayRequireLogin {
		return c.Redirect(http.StatusFound, fmt.Sprintf("/login?return=/netplay/%s/%s", game.Id, game.NetplaySessionId))
	}

	return c.Render(http.StatusOK, "netplay", pongo2.Context{
		"game":                  game,
		"user":                  user,
		"controls":              s.findNetplayControls(pctx),
		"netplay_turn_url":      s.config.TurnServerUrl,
		"netplay_turn_user":     s.config.TurnServerUser,
		"netplay_turn_password": s.config.TurnServerPassword,
		"netplay_debug":         s.config.NetplayDebug,
	})
}

func (s *Server) netplayWS(c echo.Context) error {
	pctx := c.(*PlaytimeContext)
	game := pctx.game

	isHost := false
	hostId := game.UserId
	clientId := storage.NewId()
	clientName := storage.GenerateRandomName()
	sessionId := game.NetplaySessionId

	if pctx.session != nil && len(pctx.session.UserId) != 0 {
		if pctx.session.UserId == game.UserId {
			isHost = true
			clientId = pctx.session.UserId
		}
	}

	ws, err := websocket.Accept(c.Response(), c.Request(), &websocket.AcceptOptions{CompressionMode: websocket.CompressionDisabled})
	if err != nil {
		return err
	}
	defer func() {
		if err := ws.Close(websocket.StatusNormalClosure, ""); err != nil {
			log.Debugf("unable to close ws connection if client %s in session %s: %s", clientId, sessionId, err)
		}
	}()

	user, _ := s.findContextSessionUser(pctx)
	if user != nil && len(user.Login) != 0 {
		clientName = user.Login
	}

	client := gamesession.NewClient(clientId, clientName, ws)

	s.gameSessionsMu.Lock() /////////////////////

	session := s.gameSessions.GetSession(sessionId)
	sessionNew := false
	if session == nil {
		session = gamesession.NewGameSession(game.Id, sessionId)
		sessionNew = true
	}
	if session.ClientsMaxCountReached(isHost) {
		s.gameSessionsMu.Unlock()
		s.sendWebSocketError(ws, "max client connections reached")
		log.Warnf("max client connections reached on session %s", sessionId)
		return nil
	}
	if session.GetClient(clientId) != nil {
		s.gameSessionsMu.Unlock()
		s.sendWebSocketError(ws, "client already connected")
		log.Warnf("client %s already connected in session %s", clientId, sessionId)
		return nil
	}

	session.SetClient(client)

	if sessionNew {
		s.gameSessions.SetSession(session)
	}

	s.gameSessionsMu.Unlock() ///////////////////

	session.Send(clientId, gamesession.MessageGreeting(hostId, client, s.collectNetplayCurrentSessionClients(session)))
	session.Broadcast(gamesession.MessageConnected(client))

	for {
		var incoming gamesession.MessageIncoming
		if err := wsjson.Read(context.Background(), ws, &incoming); err != nil {
			if errors.As(err, &websocket.CloseError{}) {
				log.Debugf("client %s in session %s closed ws connection", clientId, sessionId)
			} else {
				log.Warnf("client %s in session %s ws error: %s", clientId, sessionId, err)
			}
			session.RemoveClient(clientId)
			session.Broadcast(gamesession.MessageDisconnected(clientId))
			break
		}
		if len(incoming.Type) == 0 {
			continue
		}

		switch incoming.Type {

		case gamesession.MessageTypeHeartbeat:
			{
				session.SetHeartbeatReceived(clientId, true)
			}

		case gamesession.MessageTypePlayerChange:
			{
				if incoming.PlayerChange == nil {
					log.Warnf("empty player change ws message from client %s in session %s", clientId, sessionId)
					break
				}
				if !isHost && incoming.PlayerChange.ClientId != clientId {
					log.Warnf("player change ws message from non-host client %s in session %s", clientId, sessionId)
					break
				}

				changeClientId := incoming.PlayerChange.ClientId
				changePlayer := incoming.PlayerChange.Player

				if session.SetClientPlayer(changeClientId, changePlayer) {
					session.Broadcast(gamesession.MessagePlayerChanged(changeClientId, changePlayer))
				} else {
					log.Warnf("unable to player change from client %s in session %s (%s to %d)", clientId, sessionId, changeClientId, changePlayer)
				}
			}

		case gamesession.MessageTypeClientNameChange:
			{
				if incoming.NameChange == nil {
					log.Warnf("empty name change ws message from %s in session %s", clientId, sessionId)
					break
				}

				changeName := incoming.NameChange.Name

				if len([]rune(changeName)) > 32 {
					changeName = string([]rune(changeName)[0:32])
				}

				if session.SetClientName(clientId, changeName) {
					session.Broadcast(gamesession.MessageClientNameChanged(clientId, changeName))
				} else {
					log.Warnf("unable to client name change from client %s in session %s", clientId, sessionId)
				}
			}

		case gamesession.MessageTypeSignallingOffer, gamesession.MessageTypeSignallingAnswer, gamesession.MessageTypeSignallingIceCandidate:
			{
				if incoming.Signalling == nil {
					log.Warnf("empty signalling (%s) ws message from %s in session %s", incoming.Type, clientId, sessionId)
					break
				}

				destination := incoming.Signalling.ClientId
				sdp := incoming.Signalling.SDP

				if session.GetClient(destination) != nil {
					session.Send(destination, gamesession.MessageSignalling(incoming.Type, clientId, sdp))
				} else {
					log.Warnf("signalling (%s) destination client %s not connected, from client %s in session %s", incoming.Type, destination, clientId, sessionId)
				}
			}
		}
	}

	return nil
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

				session.Broadcast(gamesession.MessageDisconnected(clientId))

				continue
			}

			s.heartbeatPool.Add(func() {
				session.SetHeartbeatReceived(clientId, false)
				session.Send(clientId, gamesession.MessageHeartbeat())
			})
		}
	}
}
