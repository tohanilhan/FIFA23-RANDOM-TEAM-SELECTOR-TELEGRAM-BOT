package main

import (
	"fifa-telegram-bot/service"
	"fifa-telegram-bot/utils"
	"fifa-telegram-bot/vars"
	"log"
	"os"
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

	log.Println("Teams initialized! Total:", len(vars.Teams))
}

func main() {
	// start bot
	service.StartBot()
}
