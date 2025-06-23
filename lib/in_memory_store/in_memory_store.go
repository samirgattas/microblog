package inmemorystore

import (
	"errors"
	"log/slog"
)

var (
	ErrNotFound       = errors.New("item not found")
	ErrDuplicateEntry = errors.New("duplicate entry")
)

type Store interface {
	Save(item interface{}) error
	SaveWithID(ID int64, item interface{}) error
	Get(ID int64) (interface{}, error)
	Update(ID int64, newItem interface{}) error
	LastID() int64
	Drop()
}

type store struct {
	store map[int64]interface{}
	ID    int64
}

func NewStore() Store {
	return &store{
		store: map[int64]interface{}{},
		ID:    0,
	}
}

func (s *store) Save(item interface{}) error {
	s.ID++
	s.store[s.ID] = item
	return nil
}

func (s *store) SaveWithID(ID int64, item interface{}) error {
	if _, ok := s.store[ID]; ok {
		slog.Debug("item already exist", slog.Any("id", ID), slog.Any("method", "SaveWithID"))
		return ErrDuplicateEntry
	}
	s.store[ID] = item
	s.ID = ID
	return nil
}

func (s *store) Get(ID int64) (interface{}, error) {
	item, ok := s.store[ID]
	if !ok {
		slog.Debug("item not found", slog.Any("id", ID))
		return nil, ErrNotFound
	}
	return item, nil
}

func (s *store) Update(ID int64, newItem interface{}) error {
	_, ok := s.store[ID]
	if !ok {
		slog.Debug("item not found", slog.Any("id", ID))
		return ErrNotFound
	}
	s.store[ID] = newItem
	return nil
}

func (s *store) LastID() int64 {
	return s.ID
}

func (s *store) Drop() {
	s.store = make(map[int64]interface{})
	s.ID = 0
}
