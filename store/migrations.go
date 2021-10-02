package store

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	dynamodbPkg "github.com/fallncrlss/dictionary-app-backend/store/dynamodb"
	"github.com/pkg/errors"
)

func createWordTable(db *dynamodb.DynamoDB) error {
	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("language"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("name"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("language"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String("Word"),
	}
	if err := dynamodbPkg.CreateTable(db, params); err != nil {
		return errors.Wrap(err, "createWordTable failed")
	}
	return nil
}

func runDynamoMigrations(db *dynamodb.DynamoDB) error {
	fmt.Println("Running Migrations...")
	if err := createWordTable(db); err != nil {
		return err
	}
	fmt.Println("Migrations ran successfully!")
	return nil
}
