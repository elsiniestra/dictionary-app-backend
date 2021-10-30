package dynamodb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/fallncrlss/dictionary-app-backend/internal/lib/customerrors"
	echoLog "github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

func CreateTable(ctx context.Context, db *dynamodb.Client, params *dynamodb.CreateTableInput) error {
	_, err := db.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: params.TableName,
	})
	if err != nil {
		var rnfe *types.ResourceNotFoundException

		if errors.As(err, &rnfe) {
			_, err := db.CreateTable(ctx, params)
			if err != nil {
				return errors.Wrapf(err, "Creating table \"%s\" failed", *params.TableName)
			}

			describeParams := &dynamodb.DescribeTableInput{
				TableName: aws.String(*params.TableName),
			}

			tableExistenceWaiter := dynamodb.NewTableExistsWaiter(db)
			if err := tableExistenceWaiter.Wait(ctx, describeParams, time.Minute*2); err != nil {
				return errors.Wrapf(err, "Waiting table \"%s\" existence failed", *params.TableName)
			}

			echoLog.Debugf("Table \"%s\" successfully created!\n", *params.TableName)
		}

		echoLog.Debugf("Table \"%s\" already exists, skip...\n", *params.TableName)
	}

	return nil
}

func GetInstance(ctx context.Context, db *dynamodb.Client,
	tableName string, keyParams interface{}) (map[string]types.AttributeValue, error) {
	key, err := attributevalue.MarshalMap(keyParams)
	if err != nil {
		return nil, errors.Wrap(err, "Marshaling parameters failed")
	}

	input := dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	res, err := db.GetItem(ctx, &input)
	if err != nil {
		return nil, errors.Wrap(err, customerrors.ErrUnableFetchInstance.Error())
	}

	if res.Item == nil {
		return nil, customerrors.ErrFetchedInstanceIsNil
	}

	return res.Item, nil
}

func CreateInstance(ctx context.Context, db *dynamodb.Client, tableName string, keyParams interface{}) error {
	item, err := attributevalue.MarshalMap(keyParams)
	if err != nil {
		return errors.Wrap(err, "Marshaling parameters failed")
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = db.PutItem(ctx, input)

	if err != nil {
		return errors.Wrap(err, customerrors.ErrUnableCreateInstance.Error())
	}

	return nil
}
