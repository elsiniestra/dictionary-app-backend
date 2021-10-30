package repositories

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/fallncrlss/dictionary-app-backend/pkg/model"
)

type WordDBRepository interface {
	GetWord(name string, language string) (*model.Word, error)
	CreateWord(wordInput *model.Word) error
}

type WordWebRepository interface {
	GetWord(name string, languageCode string) (*model.Word, error)
}

type Repositories struct {
	DBWord  WordDBRepository
	WebWord WordWebRepository
}

func New(ctx context.Context, db *dynamodb.Client, client *http.Client, ODAppID string, ODAppKey string) *Repositories {
	return &Repositories{
		DBWord:  NewDynamoDBRepository(ctx, db),
		WebWord: NewODWebRepository(ctx, client, ODAppID, ODAppKey),
	}
}
