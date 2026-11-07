package autorole

import "github.com/infranoneclass/isaque-bot/internal/shared"

type AutoroleSetup interface {
	MessageID() shared.MessageID
	EmojiID() shared.EmojiID
	RoleID() shared.RoleID
}

func matchReactionToSetup(messageID shared.MessageID, emojiID shared.EmojiID, setups []AutoroleSetup) AutoroleSetup {
	for _, setup := range setups {
		if messageID == setup.MessageID() && emojiID == setup.EmojiID() {
			return setup
		}
	}
	return nil
}
