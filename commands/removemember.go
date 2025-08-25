package commands

import (
	"legends-pulse/config"
	"legends-pulse/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleMemberRemoval(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.Fields(message.Content)
	ign := strings.TrimSpace(msgSplit[1])

	err := config.RemovePlayer(ign)
	if err != nil {
		utils.SendErrorMessage(session, message.ChannelID, err)
		return
	}

	utils.SendMessage(session, message.ChannelID, ign, "Removed from the player list")

}
