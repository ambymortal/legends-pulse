package commands

import (
	"fmt"
	"legends-pulse/config"
	"legends-pulse/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

var alphabeticalList []string

func HandlePlayerList(session *discordgo.Session, message *discordgo.MessageCreate) {
	var formattedList strings.Builder

	go func() {
		loadPlayersFromConfig()

		// sort player names into alphabetical order
		c := collate.New(language.English, collate.IgnoreCase)
		c.SortStrings(alphabeticalList)

		formattedList.WriteString(fmt.Sprintf("Total Members: %v\n\n", len(alphabeticalList)))
		formattedList.WriteString("```")
		for _, ign := range alphabeticalList {
			formattedList.WriteString(fmt.Sprintf("%v\r\n", ign))
		}
		formattedList.WriteString("```")

		utils.SendMessage(session, message.ChannelID, "Player List", formattedList.String())
		alphabeticalList = nil
	}()
}

func loadPlayersFromConfig() {
	cfg := config.ParseConfig()

	for _, p := range cfg.Players {
		player := config.ConvertJsonToPlayer(p)
		alphabeticalList = append(alphabeticalList, player.Name)
	}
}
