package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/config"
	"go.uber.org/zap"
)

type AppInterface interface {
	Run()
}

type App struct {
	ServerConfig *config.ServerConfig
}

func (app *App) Run() {
	// setup logger
	logger := zap.Must(zap.NewProduction())
	if app.ServerConfig.AppEnv == config.Development {
		logger = zap.Must(zap.NewDevelopment())
	}

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println("Something went wrong while syncing the zap logger")
		}
	}(logger)

	// create the server instance
	server := &http.Server{
		Addr:         app.ServerConfig.Addr,
		ReadTimeout:  app.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: app.ServerConfig.WriteTimeout * time.Second,
		IdleTimeout:  app.ServerConfig.IdleTimeout * time.Second,
		Handler:      nil,
	}

	// start the server
	logger.Info("Starting the server...",
		zap.String("port", app.ServerConfig.Addr))

	err := server.ListenAndServe()

	if err != nil {
		logger.Error("Something went wrong while starting server")
	}
}
