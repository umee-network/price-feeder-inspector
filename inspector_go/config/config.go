package config

import (
	"encoding/json"
	"log"
	"os"
)

func ReadConfig(configFile string) (Configuration, error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("Failed to read config file. Error: %s", err)
		return Configuration{}, err
	}
	var cfg Configuration
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		log.Printf("Failed to parse config file. Error: %s", err)
		return Configuration{}, err
	}
	return cfg, nil
}
