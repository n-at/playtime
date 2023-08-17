package gamesession

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	SendTimeout = 5 * time.Second
)

type SessionStorage struct {
	sessions map[string]*GameSession
	lock     sync.RWMutex
}

func NewSessionStorage() *SessionStorage {
	storage := &SessionStorage{
		sessions: make(map[string]*GameSession),
	}
	return storage
}

func (s *SessionStorage) GetSession(sessionId string) *GameSession {
	s.lock.RLock()
	defer s.lock.RUnlock()

	session, ok := s.sessions[sessionId]
	if !ok {
		return nil
	}
	return session
}

func (s *SessionStorage) SetSession(session *GameSession) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if session == nil || len(session.GetId()) == 0 {
		return
	}
	s.sessions[session.GetId()] = session
}

func (s *SessionStorage) DeleteSession(sessionId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	session, ok := s.sessions[sessionId]
	if !ok {
		return
	}

	clients := session.GetClients()

	for _, clientId := range clients {
		if err := session.DisconnectAndRemove(clientId); err != nil {
			log.Warnf("unable to close ws connection of client %s in session %s: %s", clientId, sessionId, err)
		}
	}

	delete(s.sessions, sessionId)
}

func (s *SessionStorage) GetSessions() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var ids []string

	for sessionId := range s.sessions {
		ids = append(ids, sessionId)
	}

	return ids
}
