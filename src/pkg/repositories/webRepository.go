package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fallncrlss/dictionary-app-backend/src/internal/lib/customTypes/api"
	"github.com/fallncrlss/dictionary-app-backend/src/internal/lib/enums"
	"github.com/fallncrlss/dictionary-app-backend/src/pkg/model"

	"github.com/pkg/errors"
)

type WordWebRepository interface {
	SendRequest(url string, method string, body io.Reader) (io.ReadCloser, error)
	GetWord(name string, languageCode string) (*model.Word, error)
	Search(input string) ([]string, error)
}

type wordWebRepository struct {
	ctx    context.Context
	client *http.Client
	appID  string
	appKey string
}

func NewWordWebRepository(ctx context.Context, client *http.Client, appID string, appKey string) WordWebRepository {
	return &wordWebRepository{
		ctx:    ctx,
		client: client,
		appID:  appID,
		appKey: appKey,
	}
}

func (repo *wordWebRepository) SendRequest(url string, method string, body io.Reader) (io.ReadCloser, error) {
	request, err := http.NewRequestWithContext(repo.ctx, method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "Creating new request failed")
	}

	request.Header.Set("app_id", repo.appID)
	request.Header.Set("app_key", repo.appKey)

	response, err := repo.client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "client request to OD API failed")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	return response.Body, nil
}

func (repo *wordWebRepository) GetWord(name string, languageCode string) (*model.Word, error) {
	url := fmt.Sprintf("https://od-api.oxforddictionaries.com/api/v2/entries/%s/%s", languageCode, name)

	responseBody, err := repo.SendRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "arguments (name='%s', language='%s')", name, languageCode)
	}
	defer responseBody.Close()

	wordData := customtypes.WordDetailsAPIResponse{}
	if err := json.NewDecoder(responseBody).Decode(&wordData); err != nil {
		return nil, errors.Wrapf(err, "Decoding response from OD API request failed")
	}

	word := model.Word{
		Name:          wordData.Word,
		Phonetic:      wordData.Results[0].LexicalEntries[0].Entries[0].Pronunciations.Get().PhoneticSpelling,
		PhoneticAudio: wordData.Results[0].LexicalEntries[0].Entries[0].Pronunciations.Get().AudioFile,
		Language:      enums.GetLanguageCodes()[languageCode],
		PhrasalVerbs:  wordData.Results[0].LexicalEntries[0].PhrasalVerbs.GetTextSlice(),
		Phrases:       wordData.Results[0].LexicalEntries[0].Phrases.GetTextSlice(),
		Notes:         wordData.Results[0].LexicalEntries[0].Entries[0].Notes.GetTextSlice(),
		Senses:        wordData.Results[0].LexicalEntries[0].Entries[0].Senses.ToModel(),
	}

	return &word, nil
}

func (repo *wordWebRepository) Search(input string) ([]string, error) {
	url := "https://od-api.oxforddictionaries.com/api/v2/search/thesaurus/en?prefix=true&q=" + input
	// TODO: change request endpoint and add language param

	responseBody, err := repo.SendRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	defer responseBody.Close()

	var searchResult customtypes.SearchAPIResponse
	if err := json.NewDecoder(responseBody).Decode(&searchResult); err != nil {
		return nil, errors.Wrapf(err, "Decoding response from OD API request failed")
	}

	return searchResult.GetWordsList(), nil
}
