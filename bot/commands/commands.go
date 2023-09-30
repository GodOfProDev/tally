package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/godofprodev/tally/bot/config"
)

type Commands struct {
	app    *discordgo.Session
	config *config.BotConfig
}

func NewCommands(app *discordgo.Session, config *config.BotConfig) *Commands {
	return &Commands{app: app, config: config}
}
