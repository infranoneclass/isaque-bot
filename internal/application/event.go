package application

import (
	"github.com/infranoneclass/isaque-bot/internal/shared"
)

type EventHandler interface {
	Execute(input EventHandlerInput) error
}

type EventHandlerInput struct {
	UserID    shared.UserID
	Username  string
	ChannelID shared.ChannelID
	GuildID   shared.GuildID
	MessageID shared.MessageID
	EmojiID   shared.EmojiID
}
