package commands

import (
	"legends-pulse/config"
	"legends-pulse/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleAddCharacter(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.Fields(message.Content)

	// verify if new member is a valid character
	playerInfo, err := utils.ParseCharacterJSON(msgSplit[1])
	if err != nil {
		utils.SendErrorMessage(session, message.ChannelID, err)
		return
	}
	err2 := config.AddPlayer(playerInfo)
	if err2 != nil {
		utils.SendErrorMessage(session, message.ChannelID, err)
		return
	}

	utils.SendMessage(session, message.ChannelID, playerInfo.Name, "Successfully added to the player list!")
}
