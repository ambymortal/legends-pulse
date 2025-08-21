package pulse

import (
	"fmt"
	"legends-pulse/config"
	"legends-pulse/utils"
	"time"
)

type PlayerData struct {
	CurrentData []utils.Player
	NewData     []utils.Player
	ValidNames  []string
}

var ticker *time.Ticker

// Load all JSON player entries into currentData
// load all player names into ValidNames
func (pd *PlayerData) loadDataFromJSON() ([]utils.Player, error) {
	cfg := config.ParseConfig()

	for _, player := range cfg.Players {
		p := config.ConvertJsonToPlayer(player)
		pd.CurrentData = append(pd.CurrentData, p)
		pd.ValidNames = append(pd.ValidNames, p.Name)
	}

	return pd.CurrentData, nil
}

// generate new player data for all names included in ValidNames
func (pd *PlayerData) populateNewPlayerData() error {
	var foundError bool
	for _, name := range pd.ValidNames {
		player, err := utils.ParseCharacterJSON(name)
		if err != nil {
			fmt.Printf("Error parsing character JSON for %s: %s\n", name, err)
			foundError = true
			// go through all names before returning error
			continue
		}

		if foundError {
			return fmt.Errorf("one or more players failed to load")
		}

		pd.NewData = append(pd.NewData, player)
	}

	return nil
}

func StartMemberUpdateTask() {
	ticker = time.NewTicker(15 * time.Minute)
	data := &PlayerData{}

	go func() {
		for range ticker.C {
			// load both sets of data in prep to compare them
			data.loadDataFromJSON()
			data.populateNewPlayerData()

			// clear data so we can do it all again every 15 minutes
			data.clearData()
		}
	}()
}

func (pd *PlayerData) clearData() {
	pd.CurrentData = nil
	pd.NewData = nil
	pd.ValidNames = nil
}
