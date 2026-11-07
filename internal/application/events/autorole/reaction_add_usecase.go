package autorole

import (
	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/application/events/autorole/setups"
	"github.com/infranoneclass/isaque-bot/internal/ports"
)

type ReactionAddUseCase struct {
	discordService ports.DiscordService
	setups         []AutoroleSetup
}

func NewReactionAdd(discordService ports.DiscordService) *ReactionAddUseCase {
	autoroleSetups := []AutoroleSetup{
		setups.NewGodotAutoroleSetup(),
	}

	return &ReactionAddUseCase{
		discordService: discordService,
		setups:         autoroleSetups,
	}
}

func (u *ReactionAddUseCase) Execute(input application.EventHandlerInput) error {
	matchedSetup := matchReactionToSetup(input.MessageID, input.EmojiID, u.setups)
	if matchedSetup == nil {
		return nil
	}

	return u.discordService.AddRole(input.GuildID, input.UserID, matchedSetup.RoleID(), "User reacted to emoji")
}
