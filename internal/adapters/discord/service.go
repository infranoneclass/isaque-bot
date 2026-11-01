package discord

import (
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

func roleIsPrivileged(role discord.Role) bool {
	userCanBanMembers := (role.Permissions & discord.PermissionBanMembers) != 0
	userIsAdmin := (role.Permissions & discord.PermissionAdministrator) != 0
	return userIsAdmin || userCanBanMembers
}
