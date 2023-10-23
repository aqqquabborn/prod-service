package main

import (
	"log"
	"prod-service/app/internal/app"
	"prod-service/app/internal/config"
	"prod-service/app/pkg/logging"
)

func main() {

	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	a, err := app.NewApp(cfg, logger)

	if err != nil {
		logger.Fatal(err)

	}

	logger.Println("Running Application")

	a.Run()

}
