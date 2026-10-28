package pinghandler

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	adapter "github.com/infranoneclass/isaque-bot/internal/adapter/message_adapter"
	command "github.com/infranoneclass/isaque-bot/internal/command/ping"
)

type PingHandler struct {
	command        *command.PingCommand
	discordAdapter *adapter.MessageAdapter
}

func NewPingHandler(discordAdapter *adapter.MessageAdapter, command *command.PingCommand) *PingHandler {
	return &PingHandler{
		discordAdapter: discordAdapter,
		command:        command,
	}
}

func (mh *PingHandler) HandleCommand(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
	message := mh.command.Execute()
	return mh.discordAdapter.ToDiscordMessage(message)
}
