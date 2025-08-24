package commands

import (
	"fmt"
	"legends-pulse/config"
	"legends-pulse/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var currentList []utils.Player

func HandlePlayerList(session *discordgo.Session, message *discordgo.MessageCreate) {
	var formattedList strings.Builder

	go func() {
		loadPlayersFromConfig()

		formattedList.WriteString(fmt.Sprintf("Total Members: %v\n\n", len(currentList)))
		formattedList.WriteString("```")
		for _, char := range currentList {
			formattedList.WriteString(fmt.Sprintf("%v\r\n", char))
		}
		formattedList.WriteString("```")

		utils.SendMessage(session, message.ChannelID, "Player List", formattedList.String())
		clear()
	}()
}

func loadPlayersFromConfig() {
	cfg := config.ParseConfig()

	for _, p := range cfg.Players {
		player := config.ConvertJsonToPlayer(p)
		currentList = append(currentList, player)
	}
}

func clear() {
	currentList = nil
}
