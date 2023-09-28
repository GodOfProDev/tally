package config

import "os"

type BotConfig struct {
	Token string
}

func NewBotConfig() BotConfig {
	return BotConfig{Token: os.Getenv("BOT_TOKEN")}
}
