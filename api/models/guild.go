package models

type Guild struct {
	ServerId     int `bson:"serverId"`
	ChannelId    int `bson:"channelId"`
	CurrentCount int `bson:"currentCount"`
	HighestCount int `bson:"highestCount"`
}

func NewGuild(p *CreateGuildParams) *Guild {
	return &Guild{
		ServerId:     p.ServerId,
		ChannelId:    p.ChannelId,
		CurrentCount: 0,
		HighestCount: 0,
	}
}
