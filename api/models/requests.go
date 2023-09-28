package models

type CreateGuildParams struct {
	ServerId  int `json:"serverId"`
	ChannelId int `json:"channelId"`
}

type CreateUserParams struct {
	UserId int `json:"userId"`
}

type IncrementGuildParams struct {
	UserId int `json:"userId"`
}
