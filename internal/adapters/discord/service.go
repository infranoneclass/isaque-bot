package discord

import (
	"log"
	"slices"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/infranoneclass/isaque-bot/internal/shared"
)

type discordServiceAdapter struct {
	state *state.State
}

func NewDiscordServiceAdapter(state *state.State) *discordServiceAdapter {
	return &discordServiceAdapter{
		state: state,
	}
}

func (a *discordServiceAdapter) UserIsPrivileged(guildID shared.GuildID, userID shared.UserID) (bool, error) {
	discordUserID := discord.UserID(userID)
	discordGuildID := discord.GuildID(guildID)

	roles, err := a.state.MemberRoles(discordGuildID, discordUserID)
	if err != nil {
		return false, err
	}

	if slices.ContainsFunc(roles, roleIsPrivileged) {
		return true, nil
	}

	return false, nil
}

func (a *discordServiceAdapter) Ban(guildID shared.GuildID, userID shared.UserID, reason string, daysToDelete uint) error {
	discordUserID := discord.UserID(userID)
	discordGuildID := discord.GuildID(guildID)

	return a.state.Ban(discordGuildID, discordUserID, api.BanData{
		DeleteDays:     option.NewUint(daysToDelete),
		AuditLogReason: api.AuditLogReason(reason),
	})
}

func (a *discordServiceAdapter) Unban(guildID shared.GuildID, userID shared.UserID, reason string) error {
	discordUserID := discord.UserID(userID)
	discordGuildID := discord.GuildID(guildID)

	return a.state.Unban(discordGuildID, discordUserID, api.AuditLogReason(reason))
}

func (a *discordServiceAdapter) AddRole(guildID shared.GuildID, userID shared.UserID, roleID shared.RoleID, reason string) error {
	discordGuildID := discord.GuildID(guildID)
	discordUserID := discord.UserID(userID)
	discordRoleID := discord.RoleID(roleID)

	err := a.state.AddRole(discordGuildID, discordUserID, discordRoleID, api.AddRoleData{
		AuditLogReason: api.AuditLogReason(reason),
	})
	if err != nil {
		log.Printf("Could not add role to user %d: %s", userID, err.Error())
		return err
	}
	return nil
}

func roleIsPrivileged(role discord.Role) bool {
	userCanBanMembers := (role.Permissions & discord.PermissionBanMembers) != 0
	userIsAdmin := (role.Permissions & discord.PermissionAdministrator) != 0
	return userIsAdmin || userCanBanMembers
}
