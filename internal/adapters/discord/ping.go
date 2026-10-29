package discord

import (
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

func PingInput(cmdroute.CommandData) application.CommandInput {
	return application.CommandInput{Options: map[string]string{}}
}
