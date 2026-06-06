package main

import "github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/config"
import "github.com/ayushWeb07/AirBnb-Go-Api-Gateway/cmd/app"

func main() {
	// create the config and app instances
	cfg := &config.ServerConfig{
		Addr:         ":8080",
		ReadTimeout:  10,
		WriteTimeout: 10,
		IdleTimeout:  120,
		AppEnv:       config.Development,
	}

	serverApp := &app.App{
		ServerConfig: cfg,
	}

	// run the app
	serverApp.Run()
}
