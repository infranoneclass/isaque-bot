package autorole

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/shared"
)

func ReactionAddInputMapper(event *gateway.MessageReactionAddEvent) application.EventHandlerInput {
	return application.EventHandlerInput{
		GuildID:   shared.GuildID(event.GuildID),
		MessageID: shared.MessageID(event.MessageID),
		UserID:    shared.UserID(event.UserID),
		EmojiID:   shared.EmojiID(event.Emoji.ID),
	}
}

func Intents() []gateway.Intents {
	return []gateway.Intents{
		gateway.IntentGuildMessageReactions,
	}
}
