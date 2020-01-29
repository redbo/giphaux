package main

import (
	"log"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/redbo/giphaux"
	"github.com/redbo/giphaux/backend/sqlite"
	"github.com/redbo/giphaux/shared"
	"go.uber.org/zap"
)

func main() {
	rand.Seed(time.Now().Unix())
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	usr, err := user.Current()
	if err != nil {
		logger.Fatal("Unable to get user directory", zap.Error(err))
	}

	configFile := filepath.Join(usr.HomeDir, "giphaux.conf")
	config, err := shared.LoadConfiguration(configFile)
	if err != nil {
		logger.Fatal("Error loading config", zap.Error(err))
	}
	if config.Database == "" {
		config.Database = filepath.Join(usr.HomeDir, "giphaux-database")
	}
	if _, err := os.Stat(config.Database); err != nil && os.IsNotExist(err) {
		if err := sqlite.InitDatabase(config, logger); err != nil {
			logger.Fatal("Error creating database", zap.Error(err))
		}
	}
	giphaux.Run(config, logger)
}
