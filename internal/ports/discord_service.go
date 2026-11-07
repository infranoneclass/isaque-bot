package ports

import "github.com/infranoneclass/isaque-bot/internal/shared"

type DiscordService interface {
	Ban(guildID shared.GuildID, userID shared.UserID, reason string, daysToDelete uint) error
	Unban(guildID shared.GuildID, userID shared.UserID, reason string) error
	UserIsPrivileged(guildID shared.GuildID, userID shared.UserID) (bool, error)
	AddRole(guildID shared.GuildID, userID shared.UserID, roleID shared.RoleID, reason string) error
	RemoveRole(guildID shared.GuildID, userID shared.UserID, roleID shared.RoleID, reason string) error
}
