package ping

import (
	"context"

	"github.com/infranoneclass/isaque-bot/internal/application"
)

type UseCase struct{}

func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Definition() application.CommandDefinition {
	return application.CommandDefinition{
		Name:        "ping",
		Description: "Ping!",
	}
}

func (u *UseCase) Execute(context.Context, application.CommandInput) (application.Response, error) {
	return application.Response{Content: "Pong!"}, nil
}
