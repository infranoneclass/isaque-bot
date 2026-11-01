package honeypot

import "github.com/infranoneclass/isaque-bot/internal/shared"

type DiscordService interface {
	Ban(guildID shared.GuildID, userID shared.UserID, reason string, daysToDelete uint) error
	UserIsPrivileged(guildID shared.GuildID, userID shared.UserID) (bool, error)
}
