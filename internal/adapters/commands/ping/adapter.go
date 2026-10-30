package ping

import (
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

func InputMapper(cmdroute.CommandData) application.CommandInput {
	return application.CommandInput{Options: map[string]string{}}
}

func Intents() []gateway.Intents {
	return []gateway.Intents{
		gateway.IntentGuilds,
	}
}
