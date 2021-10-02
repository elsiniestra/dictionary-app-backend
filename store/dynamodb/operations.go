package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/fallncrlss/dictionary-app-backend/lib/customErrors"
	"github.com/pkg/errors"
)

func CreateTable(db *dynamodb.DynamoDB, params *dynamodb.CreateTableInput) error {
	_, err := db.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: params.TableName,
	})
	if awsErr, ok := err.(awserr.Error); ok {
		if awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			_, err := db.CreateTable(params)
			if err != nil {
				return err
			}

			describeParams := &dynamodb.DescribeTableInput{
				TableName: aws.String(*params.TableName),
			}
			if err := db.WaitUntilTableExists(describeParams); err != nil {
				return err
			}
			fmt.Printf("Table \"%s\" successfully created!\n", *params.TableName)
		}
	} else {
		fmt.Printf("Table \"%s\" already exists, skip...\n", *params.TableName)
	}
	return nil
}

func GetInstance(db *dynamodb.DynamoDB, tableName string, keyParams interface{}) (map[string]*dynamodb.AttributeValue, error) {
	key, err := dynamodbattribute.MarshalMap(keyParams)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}
	res, err := db.GetItem(input)
	if err != nil {
		return nil, errors.Wrap(err, customErrors.UnableFetchInstance.Error())
	}
	if res.Item == nil {
		return nil, customErrors.FetchedInstanceIsNil
	}
	return res.Item, nil
}

func CreateInstance(db *dynamodb.DynamoDB, tableName string, keyParams interface{}) error {
	item, err := dynamodbattribute.MarshalMap(keyParams)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return errors.Wrap(err, customErrors.UnableCreateInstance.Error())
	}
	return nil
}
