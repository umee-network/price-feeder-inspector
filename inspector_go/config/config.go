package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadConfig(configFile string) (Configuration, error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("Failed to read config file. Error: %s", err)
		return Configuration{}, err
	}
	var cfg Configuration
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Printf("Failed to parse config file. Error: %s", err)
		return Configuration{}, err
	}
	return cfg, nil
}
