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
	defer s.lock.RUnlock()

	p, ok := s.players[id]
	if ok {
		return p
	} else {
		return nil
	}
}

func (s *GameSession) SetPlayer(p *Player) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if p == nil || len(p.id) == 0 || len(p.name) == 0 {
		return
	}
	s.players[p.id] = p
}

func (s *GameSession) RemovePlayer(id string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.players, id)
}

func (s *GameSession) CountPlayers() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.players)
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
	s.lock.RLock()
	defer s.lock.RUnlock()

	if gamepad < 0 || gamepad > 4 {
		return nil
	}

	var playerFound *Player

	for _, player := range s.players {
		if player.GetGamepadId() == gamepad {
			playerFound = player
			break
		}
	}

	return playerFound
}

func (s *GameSession) SetPlayerGamepadId(id string, gamepad int) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(id) == 0 {
		return false
	}
	if gamepad < -1 || gamepad > 4 {
		return false
	}

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
	s.lock.RLock()
	defer s.lock.RUnlock()

	var players []string
	for playerId := range s.players {
		players = append(players, playerId)
	}
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
	defer s.lock.RUnlock()

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
