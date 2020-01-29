package shared

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configuration contains settings for the server.
type Configuration struct {
	Bind          string
	Database      string
	Verbose       bool
	UploadLimit   int64
	MaxQueryLimit int
	DomainName    string
	APIServer     bool
	WebServer     bool
}

// LoadConfiguration loads the config settings from the given .yml file.
func LoadConfiguration(configFile string) (*Configuration, error) {
	config := &Configuration{
		Bind:          "0.0.0.0:8080",
		DomainName:    "localhost:8080", // domain to use for URLs.
		Database:      "",
		Verbose:       true,
		UploadLimit:   50_000_000, // 50 MB limit on uploads by default
		MaxQueryLimit: 100,        // limit on how many results a search query can return
		APIServer:     true,
		WebServer:     true,
	}
	fp, err := os.Open(configFile)
	if err != nil { // no such file, move along with defaults.
		return config, nil
	}
	jsonData, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file: %w", err)
	}
	if err = json.Unmarshal(jsonData, config); err != nil {
		return nil, fmt.Errorf("Unable to parse json file: %w", err)
	}
	return config, nil
}
