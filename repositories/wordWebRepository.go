package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fallncrlss/dictionary-app-backend/lib/customTypes/api"
	"github.com/fallncrlss/dictionary-app-backend/lib/customerrors"
	"github.com/pkg/errors"
)

type WordWebRepository struct {
	appID  string
	appKey string
}

func NewWordWebRepository(appID string, appKey string) *WordWebRepository {
	return &WordWebRepository{
		appID:  appID,
		appKey: appKey,
	}
}

func (repo *WordWebRepository) GetWord(name string, languageCode string) (*api.WordAPIResponse, error) {
	url := fmt.Sprintf("https://od-api.oxforddictionaries.com/api/v2/entries/%s/%s", languageCode, name)
	client := http.Client{Timeout: time.Second * 2}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Creating new request failed")
	}

	request.Header.Set("app_id", repo.appID)
	request.Header.Set("app_key", repo.appKey)

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "Client request to OD API failed")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(
			customerrors.ErrIncorrectArgument, "arguments (name='%s', language='%s')", name, languageCode,
		)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	wordData := api.WordAPIResponse{}
	if err := json.NewDecoder(response.Body).Decode(&wordData); err != nil {
		return nil, errors.Wrapf(err, "Decoding response from OD API request failed")
	}

	return &wordData, nil
}
