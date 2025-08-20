package main

import (
	"fmt"
	"legends-pulse/commands"
	"legends-pulse/config"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var commandHandlers = map[string]func(session *discordgo.Session, message *discordgo.MessageCreate){
	"$character":    commands.HandleCharacterRequest,
	"$addmember":    commands.HandleAddCharacter,
	"$addcharacter": commands.HandleAddCharacter,
	"$addplayer":    commands.HandleAddCharacter,
}

func main() {
	cfg := config.ParseConfig()

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

func newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot messages
	if message.Author.ID == session.State.User.ID {
		return
	}

	// ignore messages that dont start with our command prefix
	if !strings.HasPrefix(message.Content, "$") {
		return
	}

	fields := strings.Fields(message.Content)
	if len(fields) == 0 {
		return
	}

	command := fields[0]
	handler, exists := commandHandlers[command]
	if exists {
		handler(session, message)
	}
}
