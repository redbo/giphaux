package main

import (
	"encoding/json"
	"flag"
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
	giphauxDir := filepath.Join(usr.HomeDir, ".giphaux")
	configFile := filepath.Join(giphauxDir, "giphaux.conf")
	databaseFile := filepath.Join(giphauxDir, "database")

	flag.StringVar(&configFile, "conf", configFile, "configuration file")
	flag.StringVar(&databaseFile, "db", databaseFile, "database file")
	webSrv := flag.Bool("web", false, "run only web server")
	apiSrv := flag.Bool("api", false, "run only api server")
	flag.Parse()

	config, err := shared.LoadConfiguration(configFile, databaseFile)
	if err != nil {
		logger.Fatal("Error loading config", zap.Error(err))
	}

	if *webSrv {
		config.APIServer = false
		config.WebServer = true
	} else if *apiSrv {
		config.APIServer = true
		config.WebServer = false
	}

	if len(flag.Args()) == 1 && flag.Args()[0] == "init" {
		// create a ~/.giphaux and write the default config and new database there
		err := os.Mkdir(giphauxDir, 0777)
		if err != nil && !os.IsExist(err) {
			logger.Fatal("Error creating directory", zap.Error(err))
		}
		if err := sqlite.InitDatabase(config, logger); err != nil {
			logger.Fatal("Error creating database", zap.Error(err))
		}
		fp, err := os.Create(configFile)
		if err != nil {
			logger.Fatal("Error creating config file", zap.Error(err))
		}
		defer fp.Close()
		e := json.NewEncoder(fp)
		e.SetIndent("", "    ")
		e.Encode(config)
	} else {
		giphaux.Run(config, logger)
	}
}
