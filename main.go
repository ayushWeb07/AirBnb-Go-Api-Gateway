package main

import (
	"log"

	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/config"
)
import "github.com/ayushWeb07/AirBnb-Go-Api-Gateway/cmd/app"

func main() {
	// create the config instance
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	// create the storage instance
	storage := config.CreateNewStorage()

	// create the server instance
	serverApp := &app.App{
		ServerConfig: cfg,
		Storage:      storage,
	}

	// run the app
	serverApp.Run()
}
