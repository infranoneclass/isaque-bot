package application

import "github.com/infranoneclass/isaque-bot/internal/shared"

type EventHandler interface {
	Execute(input EventHandlerInput) error
}

type EventHandlerInput struct {
	UserID    shared.UserID
	ChannelID shared.ChannelID
	GuildID   shared.GuildID
}
