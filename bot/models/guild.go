package models

type Guild struct {
	GuildId      int `json:"guildId"`
	ChannelId    int `json:"channelId"`
	CurrentCount int `json:"currentCount"`
}
