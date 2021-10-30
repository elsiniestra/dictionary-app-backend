package service

import (
	"context"

	"github.com/fallncrlss/dictionary-app-backend/pkg/model"
	"github.com/fallncrlss/dictionary-app-backend/pkg/store"
	"github.com/pkg/errors"
)

type WordService interface {
	GetWordWithDB(name string, language string) (*model.Word, error)
	GetWordWithWeb(name string, languageCode string) (*model.Word, error)
	SaveWordToDB(wordInput *model.Word) error
}

type wordService struct {
	context context.Context
	store   *store.Store
}

func NewWordService(ctx context.Context, store *store.Store) WordService {
	return &wordService{
		context: ctx,
		store:   store,
	}
}

func (ws *wordService) GetWordWithDB(name string, language string) (*model.Word, error) {
	word, err := ws.store.Repositories.DBWord.GetWord(name, language)
	if err != nil {
		return nil, errors.Wrap(err, "wordService.WordDBRepository.GetWord")
	}

	return word, nil
}

func (ws *wordService) GetWordWithWeb(name string, languageCode string) (*model.Word, error) {
	word, err := ws.store.Repositories.WebWord.GetWord(name, languageCode)
	if err != nil {
		return nil, errors.Wrap(err, "wordService.WordWebRepository.GetWord")
	}

	return word, nil
}

func (ws *wordService) SaveWordToDB(wordInput *model.Word) error {
	err := ws.store.Repositories.DBWord.CreateWord(wordInput)
	if err != nil {
		return errors.Wrap(err, "SaveWordToDB failed")
	}

	return nil
}
