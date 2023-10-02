package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/godofprodev/tally/bot/http"
	"log"
	"strconv"
)

func (h Handlers) HandleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	guild := http.NewClient().GetRequest("http://localhost:8080/v1/guilds/" + m.GuildID)

	if guild == nil {
		return
	}

	id := strconv.Itoa(guild.ChannelId)

	if m.ChannelID != id {
		return
	}

	_, err := strconv.Atoi(m.Message.Content)
	if err != nil {
		err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
		if err != nil {
			return
		}
		return
	}

	log.Print(m.Message.Content)
}
