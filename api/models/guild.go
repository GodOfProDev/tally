package models

type Guild struct {
	ServerId     int `bson:"serverId"`
	ChannelId    int `bson:"channelId"`
	CurrentCount int `bson:"currentCount"`
	HighestCount int `bson:"highestCount"`
}
