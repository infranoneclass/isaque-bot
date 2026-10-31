package messagecreate

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

func InputMapper(event *gateway.MessageCreateEvent) application.EventHandlerInput {
	return application.EventHandlerInput{
		Options: map[string]any{
			"ChannelID": event.ChannelID,
			"GuildID":   event.GuildID,
			"UserID":    event.Author.ID,
		},
	}
}

func Intents() []gateway.Intents {
	return []gateway.Intents{
		gateway.IntentGuildMessages,
	}
}
