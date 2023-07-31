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
	bot.HandleMessage("/help", func(m *tbot.Message) {
		err := c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

		_, err = c.SendMessage(m.Chat.ID, "This bot generates two random teams from the top 300 teams in FIFA 21.\n\n You can use the following commands:\n\n/generate - Generates two random teams.\n/generate -u - Updates the teams and generates two random teams.\n/help - Shows this message.\n\nMade by @tohanilhan")
	})
	bot.HandleMessage("^\\/generate(?:\\s+(-u|'-aA))?", func(m *tbot.Message) {
		err := c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

		// if -u flag is passed, update teams
		if strings.TrimLeft(m.Text, "/generate ") == "-u" || vars.Teams == nil {
			_, err = c.SendMessage(m.Chat.ID, "Updating teams...\nPlease wait.")
			if err != nil {
				return
			}
			var err error
			vars.Teams, err = utils.GetTeamsFromSoFifa()
			if err != nil {
				log.Fatal(err)
			}
		}

		generatedTeams := utils.GenerateTeams(vars.Teams)

		team1Info := fmt.Sprintf("Team 1: %s\nOVR: %d\n", generatedTeams[0].Name, generatedTeams[0].Overall)
		team2Info := fmt.Sprintf("Team 2: %s\nOVR: %d\n", generatedTeams[1].Name, generatedTeams[1].Overall)

		_, err = c.SendPhotoFile(m.Chat.ID, "images/"+generatedTeams[0].Image, tbot.OptCaption(team1Info))
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = c.SendPhotoFile(m.Chat.ID, "images/"+generatedTeams[1].Image, tbot.OptCaption(team2Info))
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Println("Teams generated for ChatID: ", m.Chat.ID)
	})
	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
