package main

import (
	"github.com/godofprodev/tally/bot/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewBotConfig()

	bot := NewBot()
	err = bot.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}

	bot.RegisterHandlers()

	err = bot.Connect()
	if err != nil {
		log.Fatal("there was an issue connecting to the bot: ", err)
	}
}
