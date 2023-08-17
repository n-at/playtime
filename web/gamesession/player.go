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
	heartbeat bool
	ws        *websocket.Conn
	lock      sync.RWMutex
}

func NewPlayer(ws *websocket.Conn) *Player {
	return &Player{
		id:        storage.NewId(),
		name:      storage.GenerateRandomName(),
		clientKey: storage.NewId(),
		gamepadId: -1,
		heartbeat: true,
		ws:        ws,
	}
}

func (p *Player) GetId() string {
	return p.id
}

func (p *Player) GetName() string {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.name
}

func (p *Player) setName(name string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(name) == 0 {
		return
	}
	p.name = name
}

func (p *Player) GetClientKey() string {
	return p.clientKey
}

func (p *Player) GetGamepadId() int {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.gamepadId
}

func (p *Player) setGamepadId(id int) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if id < -1 || id > 4 {
		return
	}
	p.gamepadId = id
}

func (p *Player) getHeartbeat() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.heartbeat
}

func (p *Player) setHeartbeat(value bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.heartbeat = value
}

func (p *Player) GetWS() *websocket.Conn {
	return p.ws
}
