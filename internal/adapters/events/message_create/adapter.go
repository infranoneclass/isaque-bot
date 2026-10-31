package messagecreate

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

func InputMapper(event *gateway.MessageCreateEvent) application.EventHandlerInput {
	return application.EventHandlerInput{
		ChannelID: uint64(event.ChannelID),
		GuildID:   uint64(event.GuildID),
		UserID:    uint64(event.Author.ID),
	}
}

func Intents() []gateway.Intents {
	return []gateway.Intents{
		gateway.IntentGuildMessages,
	}
}
