package pulse

import (
	"legends-pulse/config"
	"legends-pulse/utils"
)

var currentData []utils.Player
var validCharNames []string

// Load all JSON player entries into currentData
// load all player names into validCharNames
func LoadCurrentPlayerData() ([]utils.Player, error) {
	cfg := config.ParseConfig()

	for _, player := range cfg.Players {
		p := config.ConvertJsonToPlayer(player)
		currentData = append(currentData, p)
		validCharNames = append(validCharNames, p.Name)
	}

	return currentData, nil
}
