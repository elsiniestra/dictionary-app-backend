package repositories

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Repositories struct {
	WordDB  WordDBRepository
	WordWeb WordWebRepository
}

func New(ctx context.Context, db *dynamodb.Client, client *http.Client, ODAppID string, ODAppKey string) *Repositories {
	return &Repositories{
		WordDB:  NewWordDBRepository(ctx, db),
		WordWeb: NewWordWebRepository(ctx, client, ODAppID, ODAppKey),
	}
}
