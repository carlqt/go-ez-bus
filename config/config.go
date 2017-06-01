package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	BusKey     string
	BaseURL    string
	LTAUserKey string
}

// NewConfig returns a struct with values based on config.json file
func NewConfig() *Config {
	c := new(Config)

	workDir, _ := os.Getwd()
	configPath := filepath.Join(workDir, "config", "config.json")

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}

	return c
}
