package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Database   Database `json:"database"`
	BusKey     string   `json:"busKey"`
	BaseURL    string   `json:"baseURL"`
	LTAUserKey string   `json:"ltaUserKey"`
}

type Database struct {
	Adapter  string
	DBname   string
	SSLMode  string
	Username string
	Password string
}

// NewConfig returns a struct with values based on config.json file
func NewConfig() Config {
	c := Config{}

	workDir, _ := os.Getwd()
	configPath := filepath.Join(workDir, "config", "config.json")

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}

	return c
}
