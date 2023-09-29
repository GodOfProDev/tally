package config

import "os"

type BotConfig struct {
	Token   string
	GuildId string
}

func NewBotConfig() BotConfig {
	return BotConfig{
		Token:   os.Getenv("BOT_TOKEN"),
		GuildId: os.Getenv("GUILD_ID"),
	}
}
