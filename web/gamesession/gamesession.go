package gamesession

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"sync"
)

type GameSession struct {
	id      string
	gameId  string
	players map[string]*Player
	lock    sync.RWMutex
}

func NewGameSession(gameId, sessionId string) *GameSession {
	return &GameSession{
		id:      sessionId,
		gameId:  gameId,
		players: make(map[string]*Player),
	}
}

func (s *GameSession) GetId() string {
	return s.id
}

func (s *GameSession) GetGameId() string {
	return s.gameId
}

func (s *GameSession) GetPlayer(id string) *Player {
	s.lock.RLock()
	p, ok := s.players[id]
	if !ok {
		p = nil
	}
	s.lock.RUnlock()
	return p
}

func (s *GameSession) SetPlayer(p *Player) {
	if p == nil || len(p.id) == 0 || len(p.name) == 0 {
		return
	}
	s.lock.Lock()
	s.players[p.id] = p
	s.lock.Unlock()
}

func (s *GameSession) RemovePlayer(id string) {
	s.lock.Lock()
	delete(s.players, id)
	s.lock.Unlock()
}

func (s *GameSession) CountPlayers() int {
	s.lock.RLock()
	count := len(s.players)
	s.lock.RUnlock()
	return count
}

func (s *GameSession) SetPlayerName(id, name string) bool {
	if len(id) == 0 || len(name) == 0 {
		return false
	}

	player := s.GetPlayer(id)
	if player == nil {
		return false
	}

	player.setName(name)

	return true
}

func (s *GameSession) GetPlayerByGamepadId(gamepad int) *Player {
	if gamepad < 0 || gamepad > 4 {
		return nil
	}

	s.lock.RLock()

	var playerFound *Player

	for _, player := range s.players {
		if player.GetGamepadId() == gamepad {
			playerFound = player
			break
		}
	}

	s.lock.RUnlock()

	return playerFound
}

func (s *GameSession) SetPlayerGamepadId(id string, gamepad int) bool {
	if len(id) == 0 {
		return false
	}
	if gamepad < -1 || gamepad > 4 {
		return false
	}

	s.lock.Lock()

	ok := false
	gamepadExists := false
	var playerFound *Player

	for playerId, player := range s.players {
		if playerId == id {
			playerFound = player
		}
		if gamepad != -1 && player.gamepadId == gamepad {
			gamepadExists = true
		}
	}

	if playerFound != nil && !gamepadExists {
		playerFound.setGamepadId(gamepad)
		ok = true
	}

	s.lock.Unlock()

	return ok
}

func (s *GameSession) IsHeartbeatReceived(playerId string) bool {
	p := s.GetPlayer(playerId)
	if p == nil {
		return false
	}
	return p.getHeartbeat()
}

func (s *GameSession) SetHeartbeatReceived(playerId string, value bool) {
	p := s.GetPlayer(playerId)
	if p == nil {
		return
	}
	p.setHeartbeat(value)
}

func (s *GameSession) GetPlayers() []string {
	var players []string

	s.lock.RLock()

	for playerId := range s.players {
		players = append(players, playerId)
	}

	s.lock.RUnlock()

	return players
}

func (s *GameSession) Send(playerId string, v any) {
	player := s.GetPlayer(playerId)
	if player == nil || player.ws == nil {
		return
	}

	log.Debugf("send ws message to %s in session %s", player.id, s.id)

	ctx, cancel := context.WithTimeout(context.Background(), SendTimeout)
	defer cancel()

	if err := wsjson.Write(ctx, player.ws, v); err != nil {
		log.Warnf("unable to send ws message to %s in session %s: %s", playerId, s.id, err)
	}
}

func (s *GameSession) Broadcast(v any) {
	s.lock.RLock()

	log.Debugf("broadcast ws message in session %s", s.id)

	for _, player := range s.players {
		go func(playerId string, ws *websocket.Conn) {
			log.Debugf("broadcast ws message to player %s in session %s", playerId, s.id)

			ctx, cancel := context.WithTimeout(context.Background(), SendTimeout)
			defer cancel()

			if err := wsjson.Write(ctx, ws, v); err != nil {
				log.Warnf("unable to broadcast ws message to %s in session %s: %s", playerId, s.id, err)
			}
		}(player.id, player.ws)
	}

	s.lock.RUnlock()
}

func (s *GameSession) DisconnectAndRemove(playerId string) error {
	player := s.GetPlayer(playerId)
	if player == nil {
		return errors.New("player not found")
	}

	if player.ws != nil {
		if err := player.ws.Close(websocket.StatusNormalClosure, ""); err != nil {
			return err
		}
	}

	s.RemovePlayer(playerId)

	return nil
}
