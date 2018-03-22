package config

import (
	"os"
	"encoding/json"
)

type Config struct {
	ConnectionList []struct {
		Name     string `json:"name"`
		Adapter  string `json:"adapter"`
		Settings `json:"settings"`
	} `json:"connectionList"`
}

type Settings struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (c *Config) Load() (*Config, error) {
	configFile, err := os.Open("config/projects.json")
	defer configFile.Close()

	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(c)

	return c, nil
}