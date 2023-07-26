package storage

import "github.com/timshannon/bolthold"

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
