package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/godofprodev/tally/bot/config"
	"github.com/godofprodev/tally/bot/handlers"
	"log"
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
	command := discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Responds with pong",
		GuildID:     b.config.GuildId,
	}

	_, err := b.app.ApplicationCommandCreate(b.app.State.User.ID, b.config.GuildId, &command)
	if err != nil {
		log.Panicf("Cannot create '%v' command: %v", command.Name, err)
	}

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
	}

	b.app.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
