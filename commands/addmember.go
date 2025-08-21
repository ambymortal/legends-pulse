package commands

import (
	"legends-pulse/config"
	"legends-pulse/utils"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleAddCharacter(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.SplitAfter(message.Content, " ")

	// verify if new member is a valid character
	playerInfo, err := utils.ParseCharacterJSON(msgSplit[1])
	if err != nil {
		session.ChannelMessage(message.ChannelID, err.Error())
		return
	}
	err2 := config.AddPlayer(playerInfo)
	if err2 != nil {
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
