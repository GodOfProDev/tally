package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/godofprodev/tally/bot/http"
	"github.com/godofprodev/tally/bot/models"
	"log"
	"strconv"
)

func (c *Commands) ExecuteSetupCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	if len(options) != 1 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Invalid options provided",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})

		return
	}

	option := options[0]

	if option.Name != "count-channel" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Invalid options provided",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	gid, _ := strconv.Atoi(i.GuildID)
	cid, _ := strconv.Atoi(option.ChannelValue(c.app).ID)

	req := models.CreateGuildRequest{
		GuildId:   gid,
		ChannelId: cid,
	}

	client := http.NewClient()

	client.PostRequest(req)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Setup " + option.ChannelValue(c.app).ID,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (c *Commands) RegisterSetupCommand() {
	command := discordgo.ApplicationCommand{
		Name:        "setup",
		Description: "Setup the counting channel",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "count-channel",
				Description: "The channel it will use as it counting channel",
				ChannelTypes: []discordgo.ChannelType{
					discordgo.ChannelTypeGuildText,
				},
				Required: true,
			},
		},
	}

	_, err := c.app.ApplicationCommandCreate(c.app.State.User.ID, c.config.GuildId, &command)
	if err != nil {
		log.Panicf("Cannot create '%v' command: %v", command.Name, err)
	}
}
