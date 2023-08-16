package gamesession

import (
	"nhooyr.io/websocket"
	"playtime/storage"
	"sync"
)

type Player struct {
	id        string
	name      string
	clientKey string
	gamepadId int
	ws        *websocket.Conn
	lock      sync.RWMutex
}

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
