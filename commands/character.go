package commands

import (
	"fmt"
	"legends-pulse/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleCharacterRequest(session *discordgo.Session, message *discordgo.MessageCreate) {
	var builder strings.Builder
	msgLower := strings.ToLower(message.Content)
	msgSplit := strings.SplitAfter(msgLower, " ")

	playerInfo, err := utils.ParseCharacterJSON(msgSplit[1])
	if err != nil {
		utils.SendErrorMessage(session, message.ChannelID, err)
		return
	}
	imgUrl := fmt.Sprintf("https://maplelegends.com/api/getavatar?name=%s", playerInfo.Name)

	builder.WriteString(fmt.Sprintf("Level: %d\n", playerInfo.Level))
	builder.WriteString(fmt.Sprintf("Exp: %v\n", playerInfo.Exp))
	builder.WriteString(fmt.Sprintf("Fame: %v\n", playerInfo.Fame))
	builder.WriteString(fmt.Sprintf("Cards: %v\n", playerInfo.Cards))
	builder.WriteString(fmt.Sprintf("Job: %s\n", playerInfo.Job))
	builder.WriteString(fmt.Sprintf("Guild: %s\n", playerInfo.Guild))

	imgBuf, _ := utils.ParseChracterImage(imgUrl)

	utils.SendMessageWithImage(session, message.ChannelID, playerInfo.Name, builder.String(), imgBuf.Bytes())
}
