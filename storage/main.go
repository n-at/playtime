package storage

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/timshannon/bolthold"
	"sort"
	"time"
)

type Configuration struct {
	Path string
}

type Storage struct {
	store  *bolthold.Store
	config *Configuration
}

func New(config *Configuration) (*Storage, error) {
	s, err := bolthold.Open(config.Path, 0666, nil)
	if err != nil {
		return nil, err
	}

	return &Storage{
		store:  s,
		config: config,
	}, nil
}

func (s *Storage) Close() error {
	return s.store.Close()
}

///////////////////////////////////////////////////////////////////////////////
// User

func (s *Storage) UserSave(u User) (User, error) {
	if len(u.Login) == 0 {
		return u, errors.New("login must not be empty")
	}
	if len(u.Password) == 0 {
		return u, errors.New("password must not be empty")
	}
	if len(u.Id) == 0 {
		u.Id = NewId()
	}
	if err := s.store.Upsert(u.Id, u); err != nil {
		return u, err
	}
	return u, nil
}

func (s *Storage) UserFindById(id string) (User, error) {
	var user User
	if err := s.store.FindOne(&user, bolthold.Where(bolthold.Key).Eq(id)); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Storage) UserFindByLogin(login string) (User, error) {
	var user User
	if err := s.store.FindOne(&user, bolthold.Where("Login").Eq(login)); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Storage) UserFindBySessionId(sessionId string) (User, error) {
	session, err := s.SessionGetById(sessionId)
	if err != nil {
		return User{}, err
	}
	user, err := s.UserFindById(session.UserId)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Storage) UserFindAll() ([]User, error) {
	var users []User
	if err := s.store.Find(&users, nil); err != nil {
		return nil, err
	}
	return userSorted(users), nil
}

func (s *Storage) UserCount() (int, error) {
	return s.store.Count(User{}, nil)
}

func (s *Storage) UserDeleteById(id string) error {
	return s.store.Delete(id, User{})
}

func (s *Storage) UserEnsureExists() error {
	cnt, err := s.UserCount()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return nil
	}

	password := GenerateRandomString(10)
	encrypted, err := EncryptPassword(password)
	if err != nil {
		return err
	}

	u := User{
		Login:    "admin",
		Password: encrypted,
		Admin:    true,
		Active:   true,
	}

	_, err = s.UserSave(u)

	log.Infof(">>> ================================================")
	log.Infof(">>> created default admin user: login=%s password=%s", u.Login, password)
	log.Infof(">>> ================================================")

	return err
}

func userSorted(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Login < users[j].Login
	})
	return users
}

///////////////////////////////////////////////////////////////////////////////
// Session

func (s *Storage) SessionSave(sess Session) (Session, error) {
	if len(sess.UserId) == 0 {
		return sess, errors.New("userId must not be empty")
	}
	if len(sess.Id) == 0 {
		sess.Id = NewId()
	}
	if sess.Created.IsZero() {
		sess.Created = time.Now()
	}
	if err := s.store.Upsert(sess.Id, sess); err != nil {
		return sess, err
	}
	return sess, nil
}

func (s *Storage) SessionGetById(id string) (Session, error) {
	if len(id) == 0 {
		return Session{}, errors.New("id must not be empty")
	}
	var session Session
	if err := s.store.FindOne(&session, bolthold.Where(bolthold.Key).Eq(id)); err != nil {
		return Session{}, err
	}
	return session, nil
}

func (s *Storage) SessionDeleteById(id string) error {
	return s.store.Delete(id, Session{})
}

func (s *Storage) SessionDeleteByUserId(userId string) error {
	return s.store.DeleteMatching(Session{}, bolthold.Where("UserId").Eq(userId))
}

func (s *Storage) SessionDeleteExpired(expirationDate time.Time) error {
	return s.store.DeleteMatching(Session{}, bolthold.Where("Created").Lt(expirationDate))
}

///////////////////////////////////////////////////////////////////////////////
//Settings

func (s *Storage) SettingsSave(settings Settings) (Settings, error) {
	if len(settings.UserId) == 0 {
		return settings, errors.New("userId must not be empty")
	}
	if err := s.store.Upsert(settings.UserId, settings); err != nil {
		return settings, err
	}
	return settings, nil
}

func (s *Storage) SettingsGetByUserId(userId string) (Settings, error) {
	if len(userId) == 0 {
		return Settings{}, errors.New("userId must not be null")
	}
	var settings Settings
	err := s.store.FindOne(&settings, bolthold.Where(bolthold.Key).Eq(userId))
	if err != nil {
		if errors.Is(err, bolthold.ErrNotFound) {
			return DefaultSettings(userId), nil
		} else {
			return Settings{}, err
		}
	}
	return settings, nil
}

func (s *Storage) SettingsDeleteByUserId(userId string) error {
	return s.store.Delete(userId, Settings{})
}

///////////////////////////////////////////////////////////////////////////////
// Game

func (s *Storage) GameSave(g Game) (Game, error) {
	if len(g.UserId) == 0 {
		return g, errors.New("userId must not be empty")
	}
	if len(g.Name) == 0 {
		return g, errors.New("name must not be empty")
	}
	if len(g.Id) == 0 {
		g.Id = NewId()
	}
	if err := s.store.Upsert(g.Id, g); err != nil {
		return g, err
	}
	return g, nil
}

func (s *Storage) GameGetById(id string) (Game, error) {
	var game Game
	if err := s.store.FindOne(&game, bolthold.Where(bolthold.Key).Eq(id)); err != nil {
		return Game{}, nil
	}
	return game, nil
}

func (s *Storage) GameGetByUserId(userId string) ([]Game, error) {
	var games []Game
	if err := s.store.Find(&games, bolthold.Where("UserId").Eq(userId)); err != nil {
		return nil, err
	}
	return gameSorted(games), nil
}

func (s *Storage) GameGetByUploadBatchId(loadBatchId string) ([]Game, error) {
	lb, err := s.UploadBatchGetById(loadBatchId)
	if err != nil {
		return nil, err
	}

	var games []Game
	if err := s.store.Find(&games, bolthold.Where(bolthold.Key).In(bolthold.Slice(lb.GameIds)...)); err != nil {
		return nil, err
	}

	return gameSorted(games), nil
}

func (s *Storage) GameDeleteById(id string) error {
	return s.store.Delete(id, Game{})
}

func (s *Storage) GameDeleteByUserId(userId string) error {
	return s.store.DeleteMatching(Game{}, bolthold.Where("UserId").Eq(userId))
}

func gameSorted(games []Game) []Game {
	sort.Slice(games, func(i, j int) bool {
		gameI := &games[i]
		gameJ := &games[j]
		if gameI.Platform == gameJ.Platform {
			return gameI.Name < gameJ.Name
		}
		return gameI.Platform < gameJ.Platform
	})
	return games
}

///////////////////////////////////////////////////////////////////////////////
// SaveState

func (s *Storage) SaveStateSave(ss SaveState) (SaveState, error) {
	if len(ss.UserId) == 0 {
		return ss, errors.New("userId must not be empty")
	}
	if len(ss.GameId) == 0 {
		return ss, errors.New("gameId must not be empty")
	}

	game, err := s.GameGetById(ss.GameId)
	if err != nil {
		return ss, err
	}
	if game.UserId != ss.UserId {
		return ss, errors.New("game belongs to different user")
	}

	if len(ss.Id) == 0 {
		ss.Id = NewId()
	}
	if ss.Created.IsZero() {
		ss.Created = time.Now()
	}

	if err := s.store.Upsert(ss.Id, ss); err != nil {
		return ss, err
	}

	return ss, nil
}

func (s *Storage) SaveStateGetById(id string) (SaveState, error) {
	var ss SaveState
	if err := s.store.FindOne(&ss, bolthold.Where(bolthold.Key).Eq(id)); err != nil {
		return SaveState{}, err
	}
	return ss, nil
}

func (s *Storage) SaveStateGetByGameId(gameId string) ([]SaveState, error) {
	var ss []SaveState
	if err := s.store.Find(&ss, bolthold.Where("GameId").Eq(gameId)); err != nil {
		return nil, err
	}
	return saveStateSorted(ss), nil
}

func (s *Storage) SaveStateDeleteById(id string) error {
	return s.store.Delete(id, SaveState{})
}

func (s *Storage) SaveStateDeleteByUserId(userId string) error {
	return s.store.DeleteMatching(SaveState{}, bolthold.Where("UserId").Eq(userId))
}

func saveStateSorted(states []SaveState) []SaveState {
	sort.Slice(states, func(i, j int) bool {
		return states[i].Created.After(states[j].Created)
	})
	return states
}

///////////////////////////////////////////////////////////////////////////////
// UploadBatch

func (s *Storage) UploadBatchSave(lb UploadBatch) (UploadBatch, error) {
	if len(lb.UserId) == 0 {
		return lb, errors.New("userId must not be empty")
	}

	if len(lb.Id) == 0 {
		lb.Id = NewId()
	}
	if lb.Created.IsZero() {
		lb.Created = time.Now()
	}

	if err := s.store.Upsert(lb.Id, lb); err != nil {
		return lb, err
	}

	return lb, nil
}

func (s *Storage) UploadBatchGetById(id string) (UploadBatch, error) {
	var lb UploadBatch
	if err := s.store.FindOne(&lb, bolthold.Where(bolthold.Key).Eq(id)); err != nil {
		return UploadBatch{}, err
	}
	return lb, nil
}

func (s *Storage) UploadBatchDeleteById(id string) error {
	return s.store.Delete(id, UploadBatch{})
}

func (s *Storage) UploadBatchDeleteByUserId(userId string) error {
	return s.store.DeleteMatching(UploadBatch{}, bolthold.Where("UserId").Eq(userId))
}
