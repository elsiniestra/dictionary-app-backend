package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/pkg/errors"

	"github.com/fallncrlss/dictionary-app-backend/src/internal/cache"
	"github.com/fallncrlss/dictionary-app-backend/src/internal/config"
	"github.com/fallncrlss/dictionary-app-backend/src/internal/lib/validator"
	"github.com/fallncrlss/dictionary-app-backend/src/pkg/controller"
	"github.com/fallncrlss/dictionary-app-backend/src/pkg/service"
	"github.com/fallncrlss/dictionary-app-backend/src/pkg/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	// config
	cfg, err := config.Get()
	if err != nil {
		return err
	}

	// echo app instance
	e := echo.New()

	e.Debug = cfg.Debug
	e.Validator = validator.NewValidator()
	e.Logger.SetLevel(echoLog.INFO)

	// middlewares
	cacheClient, err := cache.New(cfg.RedisURL)
	if err != nil {
		return err
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(cacheClient.Middleware())

	// Disable Echo JSON logger in debug mode
	if cfg.Debug {
		e.Logger.SetLevel(echoLog.DEBUG)

		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}

		configJSON, err := json.Marshal(cfg)
		if err != nil {
			e.Logger.Warnf("Cannot log the config: %s", cfg)
		}

		e.Logger.Debug("Configuration:", string(configJSON))
	}

	// Init repository store
	httpClient := http.Client{Timeout: cfg.APIRequestTimeout}

	repoStore, err := store.New(ctx, e.Logger, *cfg, &httpClient)
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}

	// Init service manager
	serviceManager, err := service.NewManager(ctx, repoStore)
	if err != nil {
		return errors.Wrap(err, "manager.New failed")
	}

	// Init controllers
	wordControllers := controller.NewWordControllers(serviceManager, &e.Logger)

	// API V1
	v1 := e.Group("/api/v1")

	// Routes
	wordRoutes := v1.Group("/words")
	wordRoutes.GET("/:language/:name", wordControllers.Get)
	wordRoutes.GET("/search", wordControllers.Search)

	// Run server
	server := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
	}

	e.Logger.Fatal(e.StartServer(server))

	// nil error
	return nil
}
