package shared

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Configuration contains settings for the server.
type Configuration struct {
	Bind          string
	Database      string
	Verbose       bool
	GifsDir       string
	TempDir       string
	UploadLimit   int64
	MaxQueryLimit int
}

// LoadConfiguration loads the config settings from the given .yml file.
func LoadConfiguration(files []string) (*Configuration, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Unable to get working directory: %w", err)
	}
	config := &Configuration{
		Bind:          "0.0.0.0:8080",
		Database:      filepath.Join(cwd, "database"),
		GifsDir:       filepath.Join(cwd, "gifs"),
		Verbose:       true,
		TempDir:       os.TempDir(),
		UploadLimit:   50_000_000, // 50 MB limit on uploads by default
		MaxQueryLimit: 100,        // limit on how many results a search query can return
	}
	for _, file := range files {
		fp, err := os.Open(file)
		if err != nil { // no such file, move along.
			continue
		}
		jsonData, err := ioutil.ReadAll(fp)
		if err != nil {
			return nil, fmt.Errorf("Unable to read file: %w", err)
		}
		if err = json.Unmarshal(jsonData, config); err != nil {
			return nil, fmt.Errorf("Unable to parse json file: %w", err)
		}
	}
	return config, nil
}
