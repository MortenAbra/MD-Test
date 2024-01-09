package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Loads the config yaml file and returns a config
func LoadConfig(path string) *Config {
	configFile, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}

	return &config
}
