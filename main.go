package main

import (
	"fifa-telegram-bot/utils"
	"fifa-telegram-bot/vars"
	"fmt"
	"github.com/yanzay/tbot/v2"
	"log"
	"os"
	"strings"
	"time"
)

func init() {

	// get token from args
	vars.Token = os.Args[1]

	if vars.Token == "" {
		log.Fatal("TOKEN is empty!")
	}
	var err error
	vars.Teams, err = utils.GetTeamsFromSoFifa()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Teams fetched! Total:", len(vars.Teams))
}

func main() {

	// start bot
	bot := tbot.New(vars.Token)
	c := bot.Client()
	bot.HandleMessage("^\\/generate(?:\\s+-u)?$", func(m *tbot.Message) {
		log.Println("received /generate")
		err := c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

		// if -u flag is passed, update teams
		if strings.TrimLeft(m.Text, "/generate ") == "-u" || vars.Teams == nil {
			log.Println("Updating teams...")
			var err error
			vars.Teams, err = utils.GetTeamsFromSoFifa()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Teams updated!")
		}

		log.Println("Generating teams...")
		generatedTeams := utils.GenerateTeams(vars.Teams)
		log.Println("Teams generated!")

		// send message with teams and images
		_, err = c.SendMessage(m.Chat.ID, fmt.Sprintf("Team 1: %+v\nTeam 2: %+v", generatedTeams[0].Name, generatedTeams[1].Name))
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = c.SendPhotoFile(m.Chat.ID, "images/"+generatedTeams[0].Image, tbot.OptCaption("Team 1"))
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = c.SendPhotoFile(m.Chat.ID, "images/"+generatedTeams[1].Image, tbot.OptCaption("Team 2"))
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
