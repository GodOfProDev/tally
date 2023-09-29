package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/godofprodev/tally/bot/config"
)

type Bot struct {
	app    *discordgo.Session
	config *config.BotConfig
}

func NewBot(cfg *config.BotConfig) *Bot {
	return &Bot{config: cfg}
}

func (b *Bot) Init() error {
	app, err := discordgo.New("Bot " + b.config.Token)
	if err != nil {
		return err
	}

	b.app = app

	return nil
}

func (b *Bot) Connect() error {
	err := b.app.Open()
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) Disconnect() error {
	return b.app.Close()
}

func (b *Bot) RegisterHandlers() {

}
