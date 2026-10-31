package honeypot

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/shared"
)

func InputMapper(event *gateway.MessageCreateEvent) application.EventHandlerInput {
	return application.EventHandlerInput{
		ChannelID: shared.ChannelID(event.ChannelID),
		GuildID:   shared.GuildID(event.GuildID),
		UserID:    shared.UserID(event.Author.ID),
	}
}

func Intents() []gateway.Intents {
	return []gateway.Intents{
		gateway.IntentGuildMessages,
		gateway.IntentGuildModeration,
	}
}
