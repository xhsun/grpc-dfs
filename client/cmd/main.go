package main

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/xhsun/grpc-file-transfer/client/internal/config"
	"github.com/xhsun/grpc-file-transfer/client/internal/registry"
)

func main() {
	var config config.Config
	cleanenv.ReadEnv(&config)

	cli, err := registry.InitializeCLI(&config)
	if err != nil {
		os.Exit(2)
	}

	cli.Start()
}
