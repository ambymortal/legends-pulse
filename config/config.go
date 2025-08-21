package config

import (
	"encoding/json"
	"legends-pulse/utils"
	"log"
	"os"
)

// player specific info
type PlayerInfo struct {
	Guild  string `json:"guild"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Exp    string `json:"exp"`
	Gender string `json:"gender"`
	Job    string `json:"job"`
	Quests int    `json:"quests"`
	Cards  int    `json:"cards"`
	Donor  bool   `json:"donor"`
	Fame   int    `json:"fame"`
}

// overarching categories
type Config struct {
	Discord discordConfig `json:"discord"`
	Players []PlayerInfo  `json:"players"`
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

func saveConfig(config *Config) error {
	jsonFile, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonEncoder := json.NewEncoder(jsonFile)
	jsonEncoder.SetIndent("", "  ")

	if err := jsonEncoder.Encode(config); err != nil {
		return err
	}

	return nil
}

func AddPlayer(user utils.Player) error {
	config := ParseConfig()
	newPlayer := PlayerInfo{
		Guild:  user.Guild,
		Name:   user.Name,
		Level:  user.Level,
		Exp:    user.Exp,
		Gender: user.Gender,
		Job:    user.Job,
		Quests: user.Quests,
		Cards:  user.Cards,
		Donor:  user.Donor,
		Fame:   user.Fame,
	}

	// Append new member to existing member slice
	config.Players = append(config.Players, newPlayer)

	// Save the updated configuration to the JSON file
	if err := saveConfig(config); err != nil {
		return err
	}

	return nil
}

func ConvertJsonToPlayer(player PlayerInfo) utils.Player {
	return utils.Player{
		Guild:  player.Guild,
		Name:   player.Name,
		Level:  player.Level,
		Job:    player.Job,
		Quests: player.Quests,
		Cards:  player.Cards,
		Fame:   player.Fame,
	}
}
