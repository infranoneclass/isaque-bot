package autorole

import "github.com/infranoneclass/isaque-bot/internal/shared"

type AutoroleSetup interface {
	MessageID() shared.MessageID
	EmojiID() shared.EmojiID
	RoleID() shared.RoleID
}
