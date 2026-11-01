package bootstrap

import (
	"github.com/infranoneclass/isaque-bot/internal/adapters/discord"
)

func NewDiscordBot(config Config) *discord.Bot {
	return discord.NewBot(config.BotToken)
}
