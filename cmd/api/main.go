package main

import (
	"log"
	"net/http"
	"time"

	configPkg "github.com/fallncrlss/dictionary-app-backend/config"
	"github.com/fallncrlss/dictionary-app-backend/controller"
	loggerPkg "github.com/fallncrlss/dictionary-app-backend/logger"
	"github.com/fallncrlss/dictionary-app-backend/service"
	storePkg "github.com/fallncrlss/dictionary-app-backend/store"
	echoPkg "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// echo app instance
	echo := echoPkg.New()

	// config
	config := configPkg.Get()

	// logger
	logger := loggerPkg.Get()

	// Init repository store
	store, err := storePkg.New()
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}

	// Init service manager
	serviceManager, err := service.NewManager(store)
	if err != nil {
		return errors.Wrap(err, "manager.New failed")
	}

	// Init controllers
	wordControllers := controller.NewWordControllers(serviceManager, logger)

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
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}

	echo.Logger.Fatal(echo.StartServer(server))

	// nil error
	return nil
}
