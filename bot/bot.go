package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/godofprodev/tally/bot/config"
)

type Bot struct {
	app *discordgo.Session
}

func NewBot() *Bot {
	return &Bot{}
}

func (b *Bot) Init(cfg config.BotConfig) error {
	app, err := discordgo.New("Bot " + cfg.Token)
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
