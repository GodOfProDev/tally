package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/godofprodev/tally/bot/http"
	"strconv"
	"time"
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

	num, err := strconv.Atoi(m.Message.Content)
	if err != nil {
		err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
		if err != nil {
			return
		}
		return
	}

	if (guild.CurrentCount + 1) != num {
		err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
		if err != nil {
			return
		}
		return
	}

	avatar := m.Author.AvatarURL("")

	embed := &discordgo.MessageEmbed{
		Title:       "Current Count",
		Description: m.Message.Content,
		Footer: &discordgo.MessageEmbedFooter{
			Text: time.Now().UTC().String(),
		},
		Color: 0x03fce3,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: avatar,
		},
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}

	err = s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	if err != nil {
		return
	}

	http.NewClient().PatchRequest("http://localhost:8080/v1/guilds/"+m.GuildID+"/increment", m.Author.ID)
}
