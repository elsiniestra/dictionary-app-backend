package store

import (
	"context"
	"net/http"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/fallncrlss/dictionary-app-backend/src/internal/config"
	"github.com/fallncrlss/dictionary-app-backend/src/pkg/repositories"
)

type Store struct {
	Repositories *repositories.Repositories
}

func New(ctx context.Context, logger echo.Logger, config config.Config, client *http.Client) (*Store, error) {
	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(config.AWSRegion))
	if err != nil {
		return nil, errors.Wrap(err, "unable to load SDK config")
	}

	db := dynamodb.NewFromConfig(cfg)

	if err := runDynamoMigrations(ctx, logger, db); err != nil {
		return nil, errors.Wrap(err, "runDynamoMigrations failed")
	}

	return &Store{
		Repositories: repositories.New(ctx, db, client, config.OxfordDictionaryAppID, config.OxfordDictionaryAppKey),
	}, nil
}
