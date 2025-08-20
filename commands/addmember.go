package commands

import (
	"legends-pulse/utils"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleAddCharacter(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.SplitAfter(message.Content, " ")

	// verify if new member is a valid character
	playerInfo, _ := utils.ParseCharacterJSON(msgSplit[1])
	err := config.AddPlayer(playerInfo)
	if err != nil {
		log.Fatal(err)
	}

	session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       playerInfo.Name,
			Description: "Successfully added to the guild list",
			Color:       0x2cdaca,
		},
	})
}
