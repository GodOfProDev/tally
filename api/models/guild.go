package models

type Guild struct {
	GuildId      int `bson:"guildId" json:"guildId"`
	ChannelId    int `bson:"channelId" json:"channelId"`
	CurrentCount int `bson:"currentCount" json:"currentCount"`
	HighestCount int `bson:"highestCount" json:"highestCount"`
}

func (g *Guild) Increment() {
	g.CurrentCount++

	if g.CurrentCount > g.HighestCount {
		g.HighestCount = g.CurrentCount
	}
}

func NewGuild(p *CreateGuildParams) *Guild {
	return &Guild{
		GuildId:      p.GuildId,
		ChannelId:    p.ChannelId,
		CurrentCount: 0,
		HighestCount: 0,
	}
}
