package autorole

import (
	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/application/events/autorole/setups"
	"github.com/infranoneclass/isaque-bot/internal/ports"
)

type ReactionRemoveUseCase struct {
	discordService ports.DiscordService
	setups         []AutoroleSetup
}

func NewReactionRemove(discordService ports.DiscordService) *ReactionRemoveUseCase {
	autoroleSetups := []AutoroleSetup{
		setups.NewGodotAutoroleSetup(),
	}

	return &ReactionRemoveUseCase{
		discordService: discordService,
		setups:         autoroleSetups,
	}
}

func (u *ReactionRemoveUseCase) Execute(input application.EventHandlerInput) error {
	matchedSetup := matchReactionToSetup(input.MessageID, input.EmojiID, u.setups)
	if matchedSetup == nil {
		return nil
	}

	return u.discordService.RemoveRole(input.GuildID, input.UserID, matchedSetup.RoleID(), "User removed reaction on emoji")
}
