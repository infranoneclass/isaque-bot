package setups

import (
	messages "github.com/infranoneclass/isaque-bot/internal/application/events/autorole/setups/shared"
	"github.com/infranoneclass/isaque-bot/internal/shared"
)

type godotAutoroleSetup struct{}

const GODOT_EMOJI_ID = 746036752518610974 //1501722162942447758
const GODOT_ROLE_ID = 630917234826543105  //1497651730861789224

func NewGodotAutoroleSetup() godotAutoroleSetup {
	return godotAutoroleSetup{}
}

func (s godotAutoroleSetup) MessageID() shared.MessageID {
	return shared.MessageID(messages.ENGINE_ROLE_MESSAGE_ID)
}

func (s godotAutoroleSetup) EmojiID() shared.EmojiID {
	return shared.EmojiID(GODOT_EMOJI_ID)
}

func (s godotAutoroleSetup) RoleID() shared.RoleID {
	return shared.RoleID(GODOT_ROLE_ID)
}
