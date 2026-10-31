package honeypot

import (
	"log"
	"slices"

	"github.com/infranoneclass/isaque-bot/internal/application"
)

type UseCase struct{}

func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Execute(input application.EventHandlerInput) error {
	if !slices.Contains(GetHoneypotChannels(), input.ChannelID) {
		return nil
	}

	log.Println(input.UserID, "sent message at honeypot channel")
	return nil
}
