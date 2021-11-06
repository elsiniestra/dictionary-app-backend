package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/fallncrlss/dictionary-app-backend/src/pkg/store"
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
