package repositories

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/fallncrlss/dictionary-app-backend/lib/customErrors"
	"github.com/fallncrlss/dictionary-app-backend/model"
	dynamodbPkg "github.com/fallncrlss/dictionary-app-backend/store/dynamodb"
	"github.com/pkg/errors"
)

type WordDBRepository struct {
	db *dynamodb.DynamoDB
}

func NewWordDBRepository(db *dynamodb.DynamoDB) *WordDBRepository {
	return &WordDBRepository{db: db}
}

func (repo *WordDBRepository) GetWord(name string, language string) (*model.Word, error) {
	wordInput := model.Word{
		Name:     name,
		Language: language,
	}
	result, err := dynamodbPkg.GetInstance(repo.db, "Word", wordInput)
	if err != nil {
		return nil, err
	}

	word := model.Word{}
	err = dynamodbattribute.UnmarshalMap(result, &word)
	if err != nil {
		return nil, errors.Wrap(err, customErrors.UnableProcessInstance.Error())
	}
	return &word, nil
}

func (repo *WordDBRepository) CreateWord(wordInput *model.Word) error {
	return dynamodbPkg.CreateInstance(repo.db, "Word", wordInput)
}