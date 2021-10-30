package config

import (
	"encoding/json"
	"time"

	"github.com/kelseyhightower/envconfig"
	echoLog "github.com/labstack/gommon/log"
)

type Config struct {
	HTTPAddr               string        `envconfig:"HTTP_ADDR"`
	Debug                  bool          `envconfig:"DEBUG"`
	LogLevel               string        `envconfig:"LOG_LEVEL"`
	AWSRegion              string        `envconfig:"AWS_REGION"`
	OxfordDictionaryAppID  string        `envconfig:"OXFORD_DICTIONARY_APP_ID"`
	OxfordDictionaryAppKey string        `envconfig:"OXFORD_DICTIONARY_APP_KEY"`
	APIRequestTimeout      time.Duration `envconfig:"API_REQUEST_TIMEOUT"`
	ServerReadTimeout      time.Duration `envconfig:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout     time.Duration `envconfig:"SERVER_WRITE_TIMEOUT"`
}

func Get() *Config {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		echoLog.Fatal(err)
	}

	configBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		echoLog.Fatal(err)
	}

	echoLog.Print("Configuration:", string(configBytes))

	return &config
}
