package service

import (
	"github.com/fallncrlss/dictionary-app-backend/model"
	"github.com/fallncrlss/dictionary-app-backend/store"
	"github.com/pkg/errors"
)

type WordService struct {
	store *store.Store
}

func NewWordService(store *store.Store) *WordService {
	return &WordService{
		store: store,
	}
}

func (ws *WordService) GetWordWithDB(name string, language string) (*model.Word, error) {
	word, err := ws.store.Repositories.DBWord.GetWord(name, language)
	if err != nil {
		return nil, errors.Wrap(err, "wordService.WordDBRepository.GetWord")
	}
	return word, nil
}

func (ws *WordService) GetWordWithWeb(name string, languageCode string) (*model.Word, error) {
	wordData, err := ws.store.Repositories.WebWord.GetWord(name, languageCode)
	if err != nil {
		return nil, errors.Wrap(err, "wordService.WordWebRepository.GetWord")
	}

	word := &model.Word{
		Name:          wordData.Word,
		Phonetic:      wordData.Results[0].LexicalEntries[0].Entries[0].Pronunciations[0].PhoneticNotation,
		PhoneticAudio: wordData.Results[0].LexicalEntries[0].Entries[0].Pronunciations[0].AudioFile,
		Language:      wordData.Results[0].LexicalEntries[0].Entries[0].Pronunciations[0].Dialects[0],
		PhrasalVerbs:  wordData.Results[0].LexicalEntries[0].PhrasalVerbs.GetTextSlice(),
		Phrases:       wordData.Results[0].LexicalEntries[0].Phrases.GetTextSlice(),
		Notes:         wordData.Results[0].LexicalEntries[0].Entries[0].Notes.GetTextSlice(),
		Senses:        wordData.Results[0].LexicalEntries[0].Entries[0].Senses.ToModel(),
	}

	return word, nil
}

func (ws *WordService) SaveWordToDB(wordInput *model.Word) error {
	return ws.store.Repositories.DBWord.CreateWord(wordInput)
}
