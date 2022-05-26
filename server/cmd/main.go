package cmd

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"github.com/xhsun/grpc-file-transfer/server/internal/config"
)

func main() {
	var config config.Config
	cleanenv.ReadConfig("config/config.json", &config)
	log.WithField("config", config).Debug("Attempt to start file transfer server")

}
