package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (c *Commands) ExecutePingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Pong",
		},
	})
}

func (c *Commands) RegisterPingCommand() {

	command := discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Responds with pong",
	}

	_, err := c.app.ApplicationCommandCreate(c.app.State.User.ID, c.config.GuildId, &command)
	if err != nil {
		log.Panicf("Cannot create '%v' command: %v", command.Name, err)
	}
}
