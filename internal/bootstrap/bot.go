package bootstrap

import (
	pingadapter "github.com/infranoneclass/isaque-bot/internal/adapters/commands/ping"
	"github.com/infranoneclass/isaque-bot/internal/adapters/discord"
	pingcommand "github.com/infranoneclass/isaque-bot/internal/application/commands/ping"
)

func NewDiscordBot(config Config) *discord.Bot {
	commands := []discord.CommandBinding{
		discord.Bind(pingcommand.New(), pingadapter.InputMapper, pingadapter.Intents()),
	}

	return discord.NewBot(config.BotToken, commands)
}
