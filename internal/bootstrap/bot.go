package bootstrap

import (
	pingadapter "github.com/infranoneclass/isaque-bot/internal/adapters/commands/ping"
	"github.com/infranoneclass/isaque-bot/internal/adapters/discord"
	honeypotadapter "github.com/infranoneclass/isaque-bot/internal/adapters/events/honeypot"
	pingcommand "github.com/infranoneclass/isaque-bot/internal/application/commands/ping"
	honeypothandler "github.com/infranoneclass/isaque-bot/internal/application/events/honeypot"
)

func NewDiscordBot(config Config) *discord.Bot {
	commands := []discord.CommandBinding{
		discord.BindCommand(pingcommand.New(), pingadapter.InputMapper, pingadapter.Intents()),
	}

	eventHandlers := []discord.EventHandlerBinding{
		discord.BindEventHandler(honeypothandler.New(), honeypotadapter.InputMapper, honeypotadapter.Intents()),
	}

	return discord.NewBot(config.BotToken, commands, eventHandlers)
}
