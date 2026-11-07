package autorole

import (
	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/application/events/autorole/setups"
	"github.com/infranoneclass/isaque-bot/internal/ports"
	"github.com/infranoneclass/isaque-bot/internal/shared"
)

type AddReactionUseCase struct {
	discordService ports.DiscordService
	setups         []AutoroleSetup
}

func NewAddReaction(discordService ports.DiscordService) *AddReactionUseCase {
	autoroleSetups := []AutoroleSetup{
		setups.NewGodotAutoroleSetup(),
	}

	return &AddReactionUseCase{
		discordService: discordService,
		setups:         autoroleSetups,
	}
}

func (u *AddReactionUseCase) Execute(input application.EventHandlerInput) error {
	matchedSetup := matchReactionToSetup(input.MessageID, input.EmojiID, u.setups)
	if matchedSetup == nil {
		return nil
	}

	return u.discordService.AddRole(input.GuildID, input.UserID, matchedSetup.RoleID(), "User reacted to emoji")
}

func matchReactionToSetup(messageID shared.MessageID, emojiID shared.EmojiID, setups []AutoroleSetup) AutoroleSetup {
	for _, setup := range setups {
		if messageID == setup.MessageID() && emojiID == setup.EmojiID() {
			return setup
		}
	}
	return nil
}
