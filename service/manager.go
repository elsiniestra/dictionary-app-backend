package service

import (
	"github.com/fallncrlss/dictionary-app-backend/store"
	"github.com/pkg/errors"
)

type Manager struct {
	Word *WordService
}

func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}

	return &Manager{
		Word: NewWordService(store),
	}, nil
}
