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

	playerInfo, _ := utils.ParseCharacterJSON(msgSplit[1])

	builder.WriteString(fmt.Sprintf("Level: %d\n", playerInfo.Level))
	builder.WriteString(fmt.Sprintf("Exp: %v\n", playerInfo.Exp))
	builder.WriteString(fmt.Sprintf("Fame: %v\n", playerInfo.Fame))
	builder.WriteString(fmt.Sprintf("Cards: %v\n", playerInfo.Cards))
	builder.WriteString(fmt.Sprintf("Job: %s\n", playerInfo.Job))
	builder.WriteString(fmt.Sprintf("Guild: %s\n", playerInfo.Guild))

	session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Title:       playerInfo.Name,
		Description: builder.String(),
		Color:       0x2cdaca,
		Image: &discordgo.MessageEmbedImage{
			URL: fmt.Sprintf("https://maplelegends.com/api/getavatar?name=%s", playerInfo.Name),
		},
	})
}
