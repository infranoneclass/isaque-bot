package application

import "context"

type Command interface {
	Definition() CommandDefinition
	Execute(ctx context.Context, input CommandInput) (Response, error)
}

type CommandDefinition struct {
	Name        string
	Description string
}

type CommandInput struct {
	Options map[string]string
}

type Response struct {
	Content   string
	Ephemeral bool
}
