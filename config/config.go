package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BusKey     string
	BaseURL    string
	LTAUserKey string
}

// NewConfig returns a struct with values based on config.json file
func NewConfig() *Config {
	c := new(Config)
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}

	return c
}
