package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// discord specific info
type discordConfig struct {
	SecurityToken string
}

// overarching categories
type Config struct {
	Discord discordConfig
}

func ParseConfig() *Config {
	tomlFile := &Config{}

	if _, err := toml.DecodeFile("config.toml", tomlFile); err != nil {
		log.Fatal(err)
	}

	return tomlFile
}
