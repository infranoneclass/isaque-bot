package messagecreate

import (
	"log"

	"github.com/infranoneclass/isaque-bot/internal/application"
)

type UseCase struct{}

func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Execute(input application.EventHandlerInput) {
	log.Println(input.UserID, "sent message at channel ", input.ChannelID)
}
