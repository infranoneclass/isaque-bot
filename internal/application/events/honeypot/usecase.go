package honeypot

import (
	"fmt"
	"log"
	"slices"

	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/ports"
)

type UseCase struct {
	discordService ports.DiscordService
	configs        ports.ConfigRepository
}

func New(discordService ports.DiscordService, configs ports.ConfigRepository) *UseCase {
	return &UseCase{
		discordService: discordService,
		configs:        configs,
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

	if u.configs.GetIfHoneypotOnlyKicks() {
		err = u.discordService.Unban(input.GuildID, input.UserID, fmt.Sprintf("Usuário %s (%d) desbanido pois a feature de honeypot está em testes", input.Username, input.UserID))
		if err != nil {
			log.Printf("could not unban user %s (%d): %e", input.Username, input.UserID, err)
		}
		log.Printf("user %s (%d) was unbanned because honeypot is in test mode", input.Username, input.UserID)
	}

	return nil
}
