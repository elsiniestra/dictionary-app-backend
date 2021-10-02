package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPAddr               string `envconfig:"HTTP_ADDR"`
	LogLevel               string `envconfig:"LOG_LEVEL"`
	AWSRegion              string `envconfig:"AWS_REGION"`
	AWSEndpoint            string `envconfig:"AWS_ENDPOINT"`
	OxfordDictionaryAppID  string `envconfig:"OXFORD_DICTIONARY_APP_ID"`
	OxfordDictionaryAppKey string `envconfig:"OXFORD_DICTIONARY_APP_KEY"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Configuration:", string(configBytes))
	})

	return &config
}
