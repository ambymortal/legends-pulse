package utils

import (
	"bytes"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// helper functions for message posting

// used for all discord messages that need a character image associated
func SendMessageWithImage(session *discordgo.Session, channelID, msgTitle, msgDescription string, charImageBytes []byte) {
	session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       msgTitle,
			Description: msgDescription,
			Color:       0x2cdaca,
		},
		File: &discordgo.File{
			Name:   "output.png",
			Reader: bytes.NewReader(charImageBytes),
		},
	})
}

// default discord message
func SendMessage(session *discordgo.Session, channelID, msgTitle, msgDescription string) {
	session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       msgTitle,
			Description: msgDescription,
			Color:       0x2cdaca,
		},
	})
}

// error messaging
func SendErrorMessage(session *discordgo.Session, channelID string, err error) {
	session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Error",
			Description: fmt.Sprintf("The following error has occured: %v", err),
			Color:       0x2cdaca,
		},
	})
}
