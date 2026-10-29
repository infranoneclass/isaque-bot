package discord

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

type Bot struct {
	state       *state.State
	definitions []api.CreateCommandData
}

func NewBot(token string, commands []CommandBinding) *Bot {
	router := cmdroute.NewRouter()
	registry := NewCommandRegistry(router, NewResponsePresenter())

	botState := state.New("Bot " + token)
	botState.AddInteractionHandler(router)
	botState.AddIntents(gateway.IntentGuilds)

	return &Bot{
		state:       botState,
		definitions: registry.Register(commands),
	}
}

func (b *Bot) Run(ctx context.Context) error {
	if err := cmdroute.OverwriteCommands(b.state, b.definitions); err != nil {
		return err
	}

	return b.state.Connect(ctx)
}
