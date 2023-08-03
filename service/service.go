package service

import (
	"fifa-telegram-bot/utils"
	"fifa-telegram-bot/vars"
	"fmt"
	"github.com/yanzay/tbot/v2"
	"log"
	"time"
)

// StartBot starts the bot service
func StartBot() {
	// start bot
	bot := tbot.New(vars.Token)
	c := bot.Client()

	// handle /start command
	bot.HandleMessage("/start", func(m *tbot.Message) {
		err := c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

		_, err = c.SendMessage(m.Chat.ID, "Hello! I am a bot that generates two random teams from the top 300 teams in FIFA 23.\n\nYou can use the following commands:\n\n/random - Generates two random teams.\n/update - Updates the teams.\n/help - Shows description.\n\nMade by @tohanilhan")
		if err != nil {
			return
		}
	})

	bot.HandleMessage("/help", func(m *tbot.Message) {
		err := c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

		_, err = c.SendMessage(m.Chat.ID, "This bot generates two random teams from the top 300 teams in FIFA 23.\n\nYou can use the following commands:\n\n/random - Generates two random teams.\n/update - Updates the teams.\n/help - Shows description.\n\nMade by @tohanilhan")
	})

	bot.HandleMessage("/random", func(m *tbot.Message) {
		err := c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

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

	bot.HandleMessage("/update", func(m *tbot.Message) {

		err := c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

		_, err = c.SendMessage(m.Chat.ID, "Updating teams...\nPlease wait.")
		if err != nil {
			return
		}
		vars.Teams, err = utils.GetTeamsFromSoFifa()
		if err != nil {
			log.Fatal(err)
		}

		_, err = c.SendMessage(m.Chat.ID, "Teams updated successfully!")
		if err != nil {
			return
		}
	})

	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}

}
