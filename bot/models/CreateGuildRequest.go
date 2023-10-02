package models

type CreateGuildRequest struct {
	GuildId   int `json:"guildId"`
	ChannelId int `json:"channelId"`
}
