package main

import (
	"github.com/godofprodev/tally/bot/config"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewBotConfig()

	bot := NewBot(cfg)
	err = bot.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot.RegisterHandlers()
	err = bot.Connect()
	if err != nil {
		log.Fatal("there was an issue connecting to the bot: ", err)
	}

	bot.RegisterCommands()

	defer bot.Disconnect()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down.")
}
