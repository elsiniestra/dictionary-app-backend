package repositories

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	configPkg "github.com/fallncrlss/dictionary-app-backend/config"
)

type Repositories struct {
	DBWord  *WordDBRepository
	WebWord *WordWebRepository
}

func New(db *dynamodb.DynamoDB) *Repositories {
	config := configPkg.Get()

	return &Repositories{
		DBWord:  NewWordDBRepository(db),
		WebWord: NewWordWebRepository(config.OxfordDictionaryAppID, config.OxfordDictionaryAppKey),
	}
}
