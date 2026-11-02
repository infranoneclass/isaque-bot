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

	log.Printf("user %s (%d) sent message at honeypot channel %d", input.Username, input.UserID, input.ChannelID)

	privileged, err := u.discordService.UserIsPrivileged(input.GuildID, input.UserID)
	if err != nil {
		log.Printf("could not get user %s (%d) roles", input.Username, input.UserID)
		return err
	}

	if privileged {
		log.Printf("user %s (%d) has privileged roles, so they won't be banned", input.Username, input.UserID)
		return nil
	}

	if err = u.discordService.Ban(input.GuildID, input.UserID, "Mensagem enviada em canal proibido (honeypot)", 1); err != nil {
		log.Printf("Could not ban user %s (%d): %e", input.Username, input.UserID, err)
		return err
	}

	log.Printf("user %s (%d) was banned", input.Username, input.UserID)

	return nil
}
