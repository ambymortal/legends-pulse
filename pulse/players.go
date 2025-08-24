package pulse

import (
	"fmt"
	"legends-pulse/config"
	"legends-pulse/utils"
	"log"
	"sync"
	"time"
)

type PlayerData struct {
	CurrentData []utils.Player
	NewData     []utils.Player
	ValidNames  []string
	mu          sync.Mutex
}

var ticker *time.Ticker

// Load all JSON player entries into currentData
// load all player names into ValidNames
func (pd *PlayerData) loadDataFromJSON() error {
	cfg := config.ParseConfig()

	for _, player := range cfg.Players {
		p := config.ConvertJsonToPlayer(player)
		pd.CurrentData = append(pd.CurrentData, p)
		pd.ValidNames = append(pd.ValidNames, p.Name)
	}

	return nil
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
			data.mu.Lock()

			// load both sets of data in prep to compare them
			if err := data.loadDataFromJSON(); err != nil {
				log.Printf("error in loading current member data: %s", err)
				continue
			}
			if err := data.populateNewPlayerData(); err != nil {
				log.Printf("error in generating new member data: %s", err)
				continue
			}

			events := data.compare()
			if len(events) != 0 {
				CreatePosts(events)
			}

			// clear data so we can do it all again every 15 minutes
			data.clearData()
			data.mu.Unlock()
		}
	}()
}

// compare the CurrentData and NewData looking for differences
// when a difference between data is large enough an Event is added to the diffs var
// CARD -> every 50 cards an event achievement is generated
// FAME -> ever 50 fame an event achievement is generated
// JOB -> all job changed generate an event achievement
// LEVEL -> all level changes generate an event achievement starting at level 30
// QUEST -> every 50 quests an event achievement is generated
func (pd *PlayerData) compare() []Event {
	var diffs []Event
	for _, currentData := range pd.CurrentData {
		for _, newData := range pd.NewData {
			if currentData.Name == newData.Name {
				// cards
				if (currentData.Cards / 50) < (newData.Cards / 50) {
					diffs = append(diffs, Event{
						Name:        currentData.Name,
						Achievement: fmt.Sprintf("%s has collected %v cards!", currentData.Name, (newData.Cards/50)*50),
					})
				}
				// fame
				if (currentData.Fame / 50) < (newData.Fame / 50) {
					diffs = append(diffs, Event{
						Name:        currentData.Name,
						Achievement: fmt.Sprintf("%s has reached %v fame!", currentData.Name, (newData.Fame/50)*50),
					})
				}
				// job
				if currentData.Job != newData.Job {
					diffs = append(diffs, Event{
						Name:        currentData.Name,
						Achievement: fmt.Sprintf("%s has advanced to %s!", currentData.Name, newData.Job),
					})
				}
				// level
				if currentData.Level != newData.Level && newData.Level >= 30 {
					// only send level up posts if the player is at least level 30
					diffs = append(diffs, Event{
						Name:        currentData.Name,
						Achievement: fmt.Sprintf("%s has reached level %v!", currentData.Name, newData.Level),
					})
				}
				// quest
				if (currentData.Quests / 50) < (newData.Quests / 50) {
					// display in multiples of 50
					diffs = append(diffs, Event{
						Name:        currentData.Name,
						Achievement: fmt.Sprintf("%s has completed %v quests!", currentData.Name, (newData.Quests/50)*50),
					})
				}
			}
		}
	}

	if len(diffs) != 0 {
		fmt.Printf("Differences to be posted %s", diffs)
	}
	return diffs
}

func (pd *PlayerData) clearData() {
	pd.CurrentData = nil
	pd.NewData = nil
	pd.ValidNames = nil
}
