package honeypot

import (
	"log"
	"slices"

	"github.com/infranoneclass/isaque-bot/internal/application"
	port "github.com/infranoneclass/isaque-bot/internal/ports/honeypot"
)

type UseCase struct {
	discordService port.DiscordService
}

func New(discordService port.DiscordService) *UseCase {
	return &UseCase{
		discordService: discordService,
	}
}

func (u *UseCase) Execute(input application.EventHandlerInput) error {
	if !slices.Contains(GetHoneypotChannels(), input.ChannelID) {
		return nil
	}

	log.Println("user", input.UserID, "sent message at honeypot channel", input.ChannelID)

	privileged, err := u.discordService.UserIsPrivileged(input.GuildID, input.UserID)
	if err != nil {
		log.Println("could not get user", input.UserID, "roles")
		return err
	}

	if privileged {
		log.Println("user", input.UserID, "has privileged roles, so they won't be banned")
		return nil
	}

	// if err = u.discordService.Ban(input.GuildID, input.UserID, "Mensagem enviada em canal proibido (honeypot)", 1); err != nil {
	// 	log.Println("Could not ban user", input.UserID, ":", err.Error())
	// 	return err
	// }

	log.Println("user", input.UserID, "was banned")

	return nil
}
