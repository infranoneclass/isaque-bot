package discord

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

type CommandRegistry struct {
	router    *cmdroute.Router
	presenter *ResponsePresenter
}

type InputMapper func(cmdroute.CommandData) application.CommandInput

type CommandBinding struct {
	Command     application.Command
	InputMapper InputMapper
}

func NewCommandRegistry(router *cmdroute.Router, presenter *ResponsePresenter) *CommandRegistry {
	return &CommandRegistry{
		router:    router,
		presenter: presenter,
	}
}

func Bind(command application.Command, inputMapper InputMapper) CommandBinding {
	return CommandBinding{
		Command:     command,
		InputMapper: inputMapper,
	}
}

func (r *CommandRegistry) Register(bindings []CommandBinding) []api.CreateCommandData {
	definitions := make([]api.CreateCommandData, 0, len(bindings))

	for _, binding := range bindings {
		definition := binding.Command.Definition()
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
		input := binding.InputMapper(data)
		response, err := binding.Command.Execute(ctx, input)
		if err != nil {
			return r.presenter.Error()
		}

		return r.presenter.ToDiscord(response)
	})
}
