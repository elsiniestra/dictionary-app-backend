package main

import (
	"context"
	"github.com/fallncrlss/dictionary-app-backend/internal/lib/validator"
	"net/http"

	echoPkg "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/pkg/errors"

	"github.com/fallncrlss/dictionary-app-backend/internal/config"
	"github.com/fallncrlss/dictionary-app-backend/pkg/controller"
	"github.com/fallncrlss/dictionary-app-backend/pkg/service"
	"github.com/fallncrlss/dictionary-app-backend/pkg/store"
)

func main() {
	if err := run(); err != nil {
		echoLog.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	// config
	config := config.Get()

	// Init repository store
	httpClient := http.Client{Timeout: config.APIRequestTimeout}

	store, err := store.New(ctx, *config, &httpClient)
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}

	// Init service manager
	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		return errors.Wrap(err, "manager.New failed")
	}

	// Init controllers
	wordControllers := controller.NewWordControllers(serviceManager)

	// echo app instance
	echo := echoPkg.New()

	echo.Debug = config.Debug
	echo.Validator = validator.NewValidator()

	// Disable Echo JSON logger in debug mode
	if config.LogLevel == "DEBUG" {
		if logger, ok := echo.Logger.(*echoLog.Logger); ok {
			logger.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}

	// middlewares
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	// API V1
	v1 := echo.Group("/api/v1")

	// Routes
	wordRoutes := v1.Group("/words")
	wordRoutes.GET("/:language/:name", wordControllers.Get)

	// Run server
	server := &http.Server{
		Addr:         config.HTTPAddr,
		ReadTimeout:  config.ServerReadTimeout,
		WriteTimeout: config.ServerWriteTimeout,
	}

	echo.Logger.Fatal(echo.StartServer(server))

	// nil error
	return nil
}
