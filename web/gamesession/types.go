package gamesession

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"playtime/storage"
	"sync"
	"time"
)

type Player struct {
	id        string
	name      string
	clientKey string
	gamepadId int
	ws        *websocket.Conn
	lock      sync.RWMutex
}

type GameSession struct {
	id      string
	gameId  string
	players map[string]*Player
	lock    sync.RWMutex
}

///////////////////////////////////////////////////////////////////////////////

func NewPlayer(ws *websocket.Conn) *Player {
	return &Player{
		id:        storage.NewId(),
		name:      storage.GenerateRandomName(),
		clientKey: storage.NewId(),
		gamepadId: -1,
		ws:        ws,
	}
}

func (p *Player) GetId() string {
	return p.id
}

func (p *Player) GetName() string {
	p.lock.RLock()
	name := p.name
	p.lock.RUnlock()
	return name
}

func (p *Player) setName(name string) {
	if len(name) == 0 {
		return
	}
	p.lock.Lock()
	p.name = name
	p.lock.Unlock()
}

func (p *Player) GetClientKey() string {
	return p.clientKey
}

func (p *Player) GetGamepadId() int {
	p.lock.RLock()
	id := p.gamepadId
	p.lock.RUnlock()
	return id
}

func (p *Player) setGamepadId(id int) {
	if id < -1 || id > 4 {
		return
	}
	p.lock.Lock()
	p.gamepadId = id
	p.lock.Unlock()
}

func (p *Player) GetWS() *websocket.Conn {
	return p.ws
}

///////////////////////////////////////////////////////////////////////////////

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

func (s *GameSession) HasGamepadId(id int) bool {
	if id < 0 || id > 4 {
		return false
	}

	s.lock.RLock()

	ok := false

	for _, player := range s.players {
		if player.gamepadId == id {
			ok = true
			break
		}
	}

	s.lock.RUnlock()

	return ok
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := wsjson.Write(ctx, player.ws, v); err != nil {
		log.Warnf("unable to send ws message to %s in session %s: %s", playerId, s.id, err)
	}
}

func (s *GameSession) Broadcast(v any) {
	s.lock.RLock()

	for playerId, player := range s.players {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := wsjson.Write(ctx, player.ws, v); err != nil {
			log.Warnf("unable to broadcast ws message to %s in session %s: %s", playerId, s.id, err)
		}
	}

	s.lock.RUnlock()
}

func (s *GameSession) CloseConnection(playerId string) error {
	player := s.GetPlayer(playerId)
	if player == nil {
		return errors.New("player not found")
	}
	if player.ws == nil {
		return errors.New("player does not have ws connection")
	}

	return player.ws.Close(websocket.StatusNormalClosure, "")
}
