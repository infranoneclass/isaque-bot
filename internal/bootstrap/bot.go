package bootstrap

import (
	pingadapter "github.com/infranoneclass/isaque-bot/internal/adapters/commands/ping"
	"github.com/infranoneclass/isaque-bot/internal/adapters/discord"
	messagecreateadapter "github.com/infranoneclass/isaque-bot/internal/adapters/events/message_create"
	pingcommand "github.com/infranoneclass/isaque-bot/internal/application/commands/ping"
	messagecreatehandler "github.com/infranoneclass/isaque-bot/internal/application/events/message_create"
)

func NewDiscordBot(config Config) *discord.Bot {
	commands := []discord.CommandBinding{
		discord.BindCommand(pingcommand.New(), pingadapter.InputMapper, pingadapter.Intents()),
	}

	eventHandlers := []discord.EventHandlerBinding{
		discord.BindEventHandler(messagecreatehandler.New(), messagecreateadapter.InputMapper, messagecreateadapter.Intents()),
	}

	return discord.NewBot(config.BotToken, commands, eventHandlers)
}
