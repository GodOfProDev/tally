package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/godofprodev/tally/bot/commands"
	"github.com/godofprodev/tally/bot/config"
	"github.com/godofprodev/tally/bot/handlers"
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
	h := handlers.NewHandlers()

	b.app.AddHandler(h.HandleReady)
}

func (b *Bot) RegisterCommands() {

	cmds := commands.NewCommands(b.app, b.config)

	cmds.RegisterPingCommand()

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": cmds.ExecutePingCommand,
	}

	b.app.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
