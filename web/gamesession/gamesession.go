package gamesession

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"sync"
)

const (
	MaxClientsPerSession = 5
	MaxPlayersPerSession = 4
	PlayerSpectator      = -1
)

type GameSession struct {
	id      string
	gameId  string
	clients map[string]*Client
	lock    sync.RWMutex
}

func NewGameSession(gameId, sessionId string) *GameSession {
	return &GameSession{
		id:      sessionId,
		gameId:  gameId,
		clients: make(map[string]*Client),
	}
}

func (s *GameSession) GetId() string {
	return s.id
}

func (s *GameSession) GetGameId() string {
	return s.gameId
}

func (s *GameSession) GetClient(id string) *Client {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if p, ok := s.clients[id]; ok {
		return p
	} else {
		return nil
	}
}

func (s *GameSession) SetClient(p *Client) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if p == nil || len(p.id) == 0 || len(p.name) == 0 {
		return false
	}
	if _, ok := s.clients[p.id]; !ok {
		if len(s.clients) >= MaxClientsPerSession {
			return false
		}
	}

	s.clients[p.id] = p

	return true
}

func (s *GameSession) RemoveClient(id string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.clients, id)
}

func (s *GameSession) CountClients() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.clients)
}

func (s *GameSession) ClientsMaxCountReached(hostWantJoin bool) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	count := len(s.clients)

	if hostWantJoin {
		return count >= MaxClientsPerSession
	} else {
		//leave one free place for host
		return count >= MaxClientsPerSession-1
	}
}

func (s *GameSession) SetClientName(id, name string) bool {
	if len(id) == 0 || len(name) == 0 {
		return false
	}

	client := s.GetClient(id)
	if client == nil {
		return false
	}

	client.setName(name)

	return true
}

func (s *GameSession) SetClientPlayer(id string, player int) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(id) == 0 {
		return false
	}
	if player < 0 && player > MaxPlayersPerSession {
		if player != PlayerSpectator {
			return false
		}
	}

	ok := false
	gamepadExists := false
	var clientFound *Client

	for clientId, client := range s.clients {
		if clientId == id {
			if client.GetPlayer() == player {
				return true
			}
			clientFound = client
		}
		if player != PlayerSpectator && client.GetPlayer() == player {
			gamepadExists = true
		}
	}

	if clientFound != nil && !gamepadExists {
		clientFound.setPlayer(player)
		ok = true
	}

	return ok
}

func (s *GameSession) IsHeartbeatReceived(clientId string) bool {
	p := s.GetClient(clientId)
	if p == nil {
		return false
	}
	return p.getHeartbeat()
}

func (s *GameSession) SetHeartbeatReceived(clientId string, value bool) {
	p := s.GetClient(clientId)
	if p == nil {
		return
	}
	p.setHeartbeat(value)
}

func (s *GameSession) GetClients() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var clients []string
	for clientId := range s.clients {
		clients = append(clients, clientId)
	}
	return clients
}

///////////////////////////////////////////////////////////////////////////////

func (s *GameSession) Send(clientId string, message any) {
	client := s.GetClient(clientId)
	if client == nil || client.ws == nil {
		return
	}

	log.Debugf("send ws message to client %s in session %s", client.id, s.id)

	ctx, cancel := context.WithTimeout(context.Background(), SendTimeout)
	defer cancel()

	if err := wsjson.Write(ctx, client.ws, message); err != nil {
		log.Warnf("unable to send ws message to client %s in session %s: %s", clientId, s.id, err)
	}
}

func (s *GameSession) Broadcast(message any) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	log.Debugf("broadcast ws message in session %s", s.id)

	for _, client := range s.clients {
		go func(clientId string, ws *websocket.Conn) {
			log.Debugf("broadcast ws message to client %s in session %s", clientId, s.id)

			ctx, cancel := context.WithTimeout(context.Background(), SendTimeout)
			defer cancel()

			if err := wsjson.Write(ctx, ws, message); err != nil {
				log.Warnf("unable to broadcast ws message to client %s in session %s: %s", clientId, s.id, err)
			}
		}(client.id, client.ws)
	}
}

func (s *GameSession) DisconnectAndRemove(clientId string) error {
	client := s.GetClient(clientId)
	if client == nil {
		return errors.New("client not found")
	}

	if client.ws != nil {
		if err := client.ws.Close(websocket.StatusNormalClosure, ""); err != nil {
			return err
		}
	}

	s.RemoveClient(clientId)

	return nil
}
