package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"

	"github.com/fallncrlss/dictionary-app-backend/src/internal/lib/customerrors"
	"github.com/fallncrlss/dictionary-app-backend/src/pkg/model"
	dynamodbStore "github.com/fallncrlss/dictionary-app-backend/src/pkg/store/dynamodb"
)

type WordDBRepository interface {
	GetWord(name string, language string) (*model.Word, error)
	CreateWord(wordInput *model.Word) error
}

type dynamoDBRepository struct {
	ctx context.Context
	db  *dynamodb.Client
}

func NewWordDBRepository(ctx context.Context, db *dynamodb.Client) WordDBRepository {
	return &dynamoDBRepository{ctx: ctx, db: db}
}

func (repo *dynamoDBRepository) GetWord(name string, language string) (*model.Word, error) {
	wordInput := struct {
		Name     string `dynamodbav:"name" json:"name"`
		Language string `dynamodbav:"language" json:"language"`
	}{Name: name, Language: language}

	result, err := dynamodbStore.GetInstance(repo.ctx, repo.db, "Word", wordInput)
	if err != nil {
		return nil, errors.Wrapf(err, "Getting word \"%s\" failed", wordInput.Name)
	}

	word := model.Word{}

	err = attributevalue.UnmarshalMap(result, &word)
	if err != nil {
		return nil, errors.Wrap(err, customerrors.ErrUnableProcessInstance.Error())
	}

	return &word, nil
}

func (repo *dynamoDBRepository) CreateWord(wordInput *model.Word) error {
	err := dynamodbStore.CreateInstance(repo.ctx, repo.db, "Word", wordInput)
	if err != nil {
		return errors.Wrapf(err, "Creating word \"%s\" failed", wordInput.Name)
	}

	return nil
}
