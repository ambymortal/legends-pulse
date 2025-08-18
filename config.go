package main

import (
	"encoding/json"
	"log"
	"os"
)

// overarching categories
type Config struct {
	Discord discordConfig `json:"discord"`
}

// discord specific info
type discordConfig struct {
	SecurityToken string `json:"securityToken"`
}

func ParseConfig() *Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonDecoder := json.NewDecoder(jsonFile)
	config := &Config{}
	if err := jsonDecoder.Decode(config); err != nil {
		log.Fatal(err)
	}

	return config
}
