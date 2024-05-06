package gamesession

import (
	"nhooyr.io/websocket"
	"playtime/storage"
	"sync"
)

type Client struct {
	id        string
	name      string
	clientKey string
	player    int
	heartbeat bool
	ws        *websocket.Conn
	lock      sync.RWMutex
}

func NewClient(id, name string, ws *websocket.Conn) *Client {
	return &Client{
		id:        id,
		name:      name,
		clientKey: storage.NewId(),
		player:    PlayerSpectator,
		heartbeat: true,
		ws:        ws,
	}
}

func (p *Client) GetId() string {
	return p.id
}

func (p *Client) GetName() string {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.name
}

func (p *Client) setName(name string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(name) == 0 {
		return
	}
	p.name = name
}

func (p *Client) GetClientKey() string {
	return p.clientKey
}

func (p *Client) GetPlayer() int {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.player
}

func (p *Client) setPlayer(player int) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.player = player
}

func (p *Client) getHeartbeat() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.heartbeat
}

func (p *Client) setHeartbeat(value bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.heartbeat = value
}

func (p *Client) GetWS() *websocket.Conn {
	return p.ws
}
