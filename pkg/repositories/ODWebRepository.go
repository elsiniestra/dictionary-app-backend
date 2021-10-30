package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fallncrlss/dictionary-app-backend/internal/lib/enums"

	"github.com/fallncrlss/dictionary-app-backend/pkg/model"
	"github.com/pkg/errors"

	"github.com/fallncrlss/dictionary-app-backend/internal/lib/customTypes/api"
	"github.com/fallncrlss/dictionary-app-backend/internal/lib/customerrors"
)

type ODWebRepository struct {
	ctx    context.Context
	client *http.Client
	appID  string
	appKey string
}

func NewODWebRepository(ctx context.Context, client *http.Client, appID string, appKey string) *ODWebRepository {
	return &ODWebRepository{
		ctx:    ctx,
		client: client,
		appID:  appID,
		appKey: appKey,
	}
}

func (repo *ODWebRepository) GetWord(name string, languageCode string) (*model.Word, error) {
	url := fmt.Sprintf("https://od-api.oxforddictionaries.com/api/v2/entries/%s/%s", languageCode, name)

	request, err := http.NewRequestWithContext(repo.ctx, http.MethodGet, url, nil)
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
		return nil, errors.Wrapf(
			customerrors.ErrIncorrectArgument, "arguments (name='%s', language='%s')", name, languageCode,
		)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	wordData := customtypes.WordAPIResponse{}
	if err := json.NewDecoder(response.Body).Decode(&wordData); err != nil {
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
