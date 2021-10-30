package repositories

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/fallncrlss/dictionary-app-backend/internal/lib/customerrors"
	"github.com/fallncrlss/dictionary-app-backend/pkg/model"
	dynamodbPkg "github.com/fallncrlss/dictionary-app-backend/pkg/store/dynamodb"
	"github.com/pkg/errors"
)

type DynamoDBRepository struct {
	ctx context.Context
	db  *dynamodb.Client
}

func NewDynamoDBRepository(ctx context.Context, db *dynamodb.Client) *DynamoDBRepository {
	return &DynamoDBRepository{ctx: ctx, db: db}
}

func (repo *DynamoDBRepository) GetWord(name string, language string) (*model.Word, error) {
	wordInput := struct {
		Name     string `dynamodbav:"name" json:"name"`
		Language string `dynamodbav:"language" json:"language"`
	}{Name: name, Language: language}

	result, err := dynamodbPkg.GetInstance(repo.ctx, repo.db, "Word", wordInput)
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

func (repo *DynamoDBRepository) CreateWord(wordInput *model.Word) error {
	err := dynamodbPkg.CreateInstance(repo.ctx, repo.db, "Word", wordInput)
	if err != nil {
		return errors.Wrapf(err, "Creating word \"%s\" failed", wordInput.Name)
	}

	return nil
}
