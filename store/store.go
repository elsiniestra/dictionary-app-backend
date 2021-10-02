package store

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	configPkg "github.com/fallncrlss/dictionary-app-backend/config"
	"github.com/fallncrlss/dictionary-app-backend/repositories"
	"github.com/pkg/errors"
)

type Store struct {
	Repositories *repositories.Repositories
}

func New() (*Store, error) {
	log.Println("Running DynamoDB migrations...")

	config := configPkg.Get()

	awsConfig := &aws.Config{
		Region:   aws.String(config.AWSRegion),
		Endpoint: aws.String(config.AWSEndpoint),
	}
	awsSession := session.Must(session.NewSession(awsConfig))
	db := dynamodb.New(awsSession)

	if err := runDynamoMigrations(db); err != nil {
		return nil, errors.Wrap(err, "runDynamoMigrations failed")
	}

	return &Store{
		Repositories: repositories.New(db),
	}, nil
}
