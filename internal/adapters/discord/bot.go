package discord

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

type Bot struct {
	state    *state.State
	commands []api.CreateCommandData
}

func NewBot(token string, commands []CommandBinding, eventHandlers []EventHandlerBinding) *Bot {
	botState := state.New("Bot " + token)

	router := cmdroute.NewRouter()
	commandRegistry := NewCommandRegistry(router, NewResponsePresenter())
	botState.AddInteractionHandler(router)

	eventHandlerRegistry := NewEventHandlerRegistry(botState)
	eventHandlerRegistry.Register(eventHandlers)

	var intents gateway.Intents
	for _, command := range commands {
		intents |= MergeIntents(command.Intents()...)
	}
	for _, handler := range eventHandlers {
		intents |= MergeIntents(handler.Intents()...)
	}
	botState.AddIntents(intents)

	return &Bot{
		state:    botState,
		commands: commandRegistry.Register(commands),
	}
}

func (b *Bot) Run(ctx context.Context) error {
	if err := cmdroute.OverwriteCommands(b.state, b.commands); err != nil {
		return err
	}

	return b.state.Connect(ctx)
}
