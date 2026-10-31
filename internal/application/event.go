package application

type EventHandler interface {
	Execute(input EventHandlerInput)
}

type EventHandlerInput struct {
	UserID    uint64
	ChannelID uint64
	GuildID   uint64
}
