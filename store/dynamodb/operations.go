package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/fallncrlss/dictionary-app-backend/lib/customerrors"
	loggerPkg "github.com/fallncrlss/dictionary-app-backend/logger"
	"github.com/pkg/errors"
)

func CreateTable(db *dynamodb.DynamoDB, params *dynamodb.CreateTableInput) error {
	logger := loggerPkg.Get()

	_, err := db.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: params.TableName,
	})

	awsErr, ok := err.(awserr.Error)
	if !ok {
		logger.Debug().Msgf("Table \"%s\" already exists, skip...\n", *params.TableName)

		return nil
	}

	if awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
		_, err := db.CreateTable(params)
		if err != nil {
			return errors.Wrapf(err, "Creating table \"%s\" failed", *params.TableName)
		}

		describeParams := &dynamodb.DescribeTableInput{
			TableName: aws.String(*params.TableName),
		}
		if err := db.WaitUntilTableExists(describeParams); err != nil {
			return errors.Wrapf(err, "Waiting table \"%s\" existence failed", *params.TableName)
		}

		logger.Debug().Msgf("Table \"%s\" successfully created!\n", *params.TableName)
	}

	return nil
}

func GetInstance(db *dynamodb.DynamoDB, tableName string, keyParams interface{}) (map[string]*dynamodb.AttributeValue, error) {
	key, err := dynamodbattribute.MarshalMap(keyParams)
	if err != nil {
		return nil, errors.Wrap(err, "Marshaling parameters failed")
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	res, err := db.GetItem(input)
	if err != nil {
		return nil, errors.Wrap(err, customerrors.ErrUnableFetchInstance.Error())
	}

	if res.Item == nil {
		return nil, customerrors.ErrFetchedInstanceIsNil
	}

	return res.Item, nil
}

func CreateInstance(db *dynamodb.DynamoDB, tableName string, keyParams interface{}) error {
	item, err := dynamodbattribute.MarshalMap(keyParams)
	if err != nil {
		return errors.Wrap(err, "Marshaling parameters failed")
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return errors.Wrap(err, customerrors.ErrUnableCreateInstance.Error())
	}

	return nil
}
