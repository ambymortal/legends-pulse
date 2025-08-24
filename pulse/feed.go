package pulse

import (
	"fmt"
	"legends-pulse/utils"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Event struct {
	Name        string
	Achievement string
}

var feedChannel string
var s *discordgo.Session
var active bool

func SetFeedChannel(session *discordgo.Session, message *discordgo.MessageCreate) {
	//always have the player notifications post in same designated channel
	feedChannel = message.ChannelID
	s = session
	active = true

	utils.SendMessage(session, message.ChannelID, "Crossroads Alliance Updates", "Updates have now been turned on")
}

func CreatePosts(events []Event) {
	if !active {
		// make sure feed channel has been established before post creation
		return
	}

	for _, event := range events {
		charUrl := fmt.Sprintf("https://maplelegends.com/api/getavatar?name=%s", event.Name)

		imgBuf, err := utils.ParseChracterImage(charUrl)
		if err != nil {
			log.Println("Error occurred:", err)
			continue
		}

		utils.SendMessageWithImage(s, feedChannel, "Player Update", event.Achievement, imgBuf.Bytes())
	}
}
