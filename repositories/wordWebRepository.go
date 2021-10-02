package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/fallncrlss/dictionary-app-backend/lib/customErrors"
	"github.com/fallncrlss/dictionary-app-backend/lib/customTypes/api"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type WordWebRepository struct {
	appId  string
	appKey string
}

func NewWordWebRepository(appId string, appKey string) *WordWebRepository {
	return &WordWebRepository{
		appId:  appId,
		appKey: appKey,
	}
}

func (repo *WordWebRepository) GetWord(name string, languageCode string) (*api.WordApiResponse, error) {
	url := fmt.Sprintf("https://od-api.oxforddictionaries.com/api/v2/entries/%s/%s", languageCode, name)
	client := http.Client{Timeout: time.Second * 2}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("app_id", repo.appId)
	request.Header.Set("app_key", repo.appKey)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(customErrors.IncorrectArgument, "arguments (name='%s', language='%s')", name, languageCode)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	wordData := api.WordApiResponse{}
	if err := json.NewDecoder(response.Body).Decode(&wordData); err != nil {
		return nil, err
	}
	return &wordData, nil
}
