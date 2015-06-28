package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
)

//Config is to store instance specific configuration information.
type Config struct {
	First *index.Year
	Last  *index.Year
}

func pathToConfig(kd kaydir.KayDir) string {
	return filepath.Join(string(kd), "config.json")
}

//Get creates an in memory Config from the config.json file of the .kay directory.  If config does not exist an empty one is returned.
func Get(kd kaydir.KayDir) (*Config, error) {
	path := pathToConfig(kd)

	_, statErr := os.Stat(path)
	if os.IsNotExist(statErr) {
		return &Config{}, nil
	}

	if statErr != nil {
		return nil, statErr
	}

	file, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}

	var config Config
	jsonErr := json.Unmarshal(file, &config)

	if jsonErr != nil {
		return nil, jsonErr
	}

	return &config, nil
}

//Set writes the in memory Config to the config.json file in the .kay directory.
func Set(kd kaydir.KayDir, config *Config) error {
	path := pathToConfig(kd)

	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, bytes, 0777)
}
