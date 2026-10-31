package discord

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

type CommandRegistry struct {
	router    *cmdroute.Router
	presenter *ResponsePresenter
}

type CommandInputMapper func(cmdroute.CommandData) application.CommandInput

type CommandBinding struct {
	command     application.Command
	inputMapper CommandInputMapper
	intents     []gateway.Intents
}

func NewCommandRegistry(router *cmdroute.Router, presenter *ResponsePresenter) *CommandRegistry {
	return &CommandRegistry{
		router:    router,
		presenter: presenter,
	}
}

func (c CommandBinding) Intents() []gateway.Intents {
	return c.intents
}

func BindCommand(command application.Command, inputMapper CommandInputMapper, intents []gateway.Intents) CommandBinding {
	return CommandBinding{
		command:     command,
		inputMapper: inputMapper,
		intents:     intents,
	}
}

func (r *CommandRegistry) Register(bindings []CommandBinding) []api.CreateCommandData {
	definitions := make([]api.CreateCommandData, 0, len(bindings))

	for _, binding := range bindings {
		definition := binding.command.Definition()
		definitions = append(definitions, api.CreateCommandData{
			Name:        definition.Name,
			Description: definition.Description,
		})

		r.register(binding, definition.Name)
	}

	return definitions
}

func (r *CommandRegistry) register(binding CommandBinding, name string) {
	r.router.AddFunc(name, func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		input := binding.inputMapper(data)
		response, err := binding.command.Execute(ctx, input)
		if err != nil {
			return r.presenter.Error()
		}

		return r.presenter.ToDiscord(response)
	})
}
