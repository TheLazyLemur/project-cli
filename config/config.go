package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	StoreDirectory string `json:"storeDirectory"`
	DefaultEditor  string `json:"defaultEditor"`
}

func GetConfig() (*Config, error) {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	c := new(Config)
	b, err := os.ReadFile(dirname + "/.config/project-cli/config.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, &c)
	return c, err
}
