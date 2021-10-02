package store

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	loggerPkg "github.com/fallncrlss/dictionary-app-backend/logger"
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
	logger := loggerPkg.Get()

	logger.Debug().Msg("Running Migrations...")

	if err := createWordTable(db); err != nil {
		return err
	}

	logger.Debug().Msg("Migrations ran successfully!")

	return nil
}
