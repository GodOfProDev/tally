package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (c *Commands) ExecuteSetupCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Setup",
		},
	})
}

func (c *Commands) RegisterSetupCommand() {

	command := discordgo.ApplicationCommand{
		Name:        "setup",
		Description: "Responds with setup",
	}

	_, err := c.app.ApplicationCommandCreate(c.app.State.User.ID, c.config.GuildId, &command)
	if err != nil {
		log.Panicf("Cannot create '%v' command: %v", command.Name, err)
	}
}
