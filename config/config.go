package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	StoreDirectory string `json:"storeDirectory"`
	DefaultEditor  string `json:"defaultEditor"`
}

func GetConfig() (*Config, error) {
	c := new(Config)
	b, err := os.ReadFile("/home/dan/.config/project-cli/config.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, &c)
	return c, err
}
