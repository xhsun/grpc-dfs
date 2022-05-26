package main

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
	"github.com/xhsun/grpc-file-transfer/server/internal/registry"
)

func main() {
	var config config.Config
	cleanenv.ReadConfig("config/config.json", &config)
	log.WithField("config", config).Debug("Attempt to start file transfer server")

	// Intialize services
	appServers, err := registry.InitializeServer(&config)
	if err != nil {
		log.WithError(err).Error("Failed to initialize file transfer server")
		os.Exit(2)
	}

	// Start services
	err = appServers.Start()
	if err != nil {
		log.WithError(err).Error("Failed to start file transfer server")
		os.Exit(2)
	}
}
