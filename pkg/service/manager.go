package service

import (
	"context"

	"github.com/fallncrlss/dictionary-app-backend/pkg/store"
	"github.com/pkg/errors"
)

type Manager struct {
	Word WordService
}

func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}

	return &Manager{
		Word: NewWordService(ctx, store),
	}, nil
}
