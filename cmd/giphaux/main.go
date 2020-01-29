package main

import (
	"log"
	"math/rand"
	"os/user"
	"path/filepath"
	"time"

	"github.com/redbo/giphaux"
	"github.com/redbo/giphaux/shared"
)

func main() {
	rand.Seed(time.Now().Unix())

	// We look for giphaux.conf in the following order: ./giphau.conf (current working directory),
	// 		~/giphaux.conf (home directory), /etc/giphaux.conf
	configFiles := []string{"giphaux.conf"}
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get home directory: %v", err)
	} else {
		configFiles = append(configFiles, filepath.Join(usr.HomeDir, "giphaux.conf"))
	}
	configFiles = append(configFiles, "/etc/giphaux.conf")

	config, err := shared.LoadConfiguration(configFiles)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	giphaux.Run(config)
}
