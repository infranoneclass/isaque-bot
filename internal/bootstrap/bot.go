package bootstrap

import (
	"github.com/infranoneclass/isaque-bot/internal/adapters/discord"
	"github.com/infranoneclass/isaque-bot/internal/application/ping"
)

func NewDiscordBot(config Config) *discord.Bot {
	commands := []discord.CommandBinding{
		discord.Bind(ping.New(), discord.PingInput),
	}

	return discord.NewBot(config.BotToken, commands)
}
