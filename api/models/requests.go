package models

type CreateGuildParams struct {
	GuildId   int `json:"guildId"`
	ChannelId int `json:"channelId"`
}

type CreateUserParams struct {
	UserId int `json:"userId"`
}

type IncrementGuildParams struct {
	UserId int `json:"userId"`
}
