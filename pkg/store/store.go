package store

import (
	"context"
	"net/http"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/fallncrlss/dictionary-app-backend/internal/config"
	"github.com/fallncrlss/dictionary-app-backend/pkg/repositories"
	echoLog "github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

type Store struct {
	Repositories *repositories.Repositories
}

func New(ctx context.Context, config config.Config, client *http.Client) (*Store, error) {
	echoLog.Debug("Running DynamoDB migrations...")

	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(config.AWSRegion))
	if err != nil {
		return nil, errors.Wrap(err, "unable to load SDK config")
	}

	db := dynamodb.NewFromConfig(cfg)
	if err := runDynamoMigrations(ctx, db); err != nil {
		return nil, errors.Wrap(err, "runDynamoMigrations failed")
	}

	return &Store{
		Repositories: repositories.New(ctx, db, client, config.OxfordDictionaryAppID, config.OxfordDictionaryAppKey),
	}, nil
}
