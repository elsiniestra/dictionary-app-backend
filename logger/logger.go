package logger

import (
	"os"
	"strings"
	"sync"

	"github.com/fallncrlss/dictionary-app-backend/config"
	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

var (
	logger Logger
	once   sync.Once
)

func Get() *Logger {
	once.Do(func() {
		zeroLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		cfg := config.Get()
		// Set proper loglevel based on config
		switch strings.ToLower(cfg.LogLevel) {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn", "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "err", "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		logger = Logger{&zeroLogger}
	})

	return &logger
}
