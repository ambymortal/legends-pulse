package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func Init() {
	cfg := ParseConfig()

	discord, err := discordgo.New("Bot " + cfg.Discord.SecurityToken)
	if err != nil {
		log.Fatal(err)
	}

	// event handlers
	discord.AddHandler(newMessage)

	// Open websocket
	discord.Open()

	// Run until process is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//close websocket
	defer discord.Close()
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// test command
	switch {
	case message.Content == "~test":
		discord.ChannelMessageSend(message.ChannelID, "response")
	}
}
