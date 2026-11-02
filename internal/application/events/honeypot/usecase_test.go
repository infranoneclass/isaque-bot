package honeypot_test

import (
	"errors"
	"testing"

	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/application/events/honeypot"
	"github.com/infranoneclass/isaque-bot/internal/ports/honeypot/mocks"
	"github.com/infranoneclass/isaque-bot/internal/shared"
)

func TestExecute(t *testing.T) {
	serviceErr := errors.New("discord service unavailable")
	banErr := errors.New("could not ban user")
	tests := []struct {
		name    string
		input   application.EventHandlerInput
		setup   func(*mocks.MockDiscordService, application.EventHandlerInput)
		wantErr error
	}{
		{
			name: "ignores messages outside honeypot channels",
			input: application.EventHandlerInput{
				GuildID:   shared.GuildID(1),
				UserID:    shared.UserID(2),
				ChannelID: shared.ChannelID(3),
			},
		},
		{
			name: "returns the privilege lookup error",
			input: application.EventHandlerInput{
				GuildID:   shared.GuildID(1),
				UserID:    shared.UserID(2),
				ChannelID: honeypot.OBJ_ISAQUE_CHANNEL_ID,
			},
			setup: func(discordService *mocks.MockDiscordService, input application.EventHandlerInput) {
				discordService.EXPECT().UserIsPrivileged(input.GuildID, input.UserID).Return(false, serviceErr)
			},
			wantErr: serviceErr,
		},
		{
			name: "allows privileged users in honeypot channels",
			input: application.EventHandlerInput{
				GuildID:   shared.GuildID(1),
				UserID:    shared.UserID(2),
				ChannelID: honeypot.NONECLASS_CHANNEL_ID,
			},
			setup: func(discordService *mocks.MockDiscordService, input application.EventHandlerInput) {
				discordService.EXPECT().UserIsPrivileged(input.GuildID, input.UserID).Return(true, nil)
			},
		},
		{
			name: "bans unprivileged users in honeypot channels",
			input: application.EventHandlerInput{
				GuildID:   shared.GuildID(1),
				UserID:    shared.UserID(2),
				ChannelID: honeypot.OBJ_ISAQUE_CHANNEL_ID,
			},
			setup: func(discordService *mocks.MockDiscordService, input application.EventHandlerInput) {
				discordService.EXPECT().UserIsPrivileged(input.GuildID, input.UserID).Return(false, nil)
				discordService.EXPECT().Ban(input.GuildID, input.UserID, "Mensagem enviada em canal proibido (honeypot)", uint(1)).Return(nil)
			},
		},
		{
			name: "returns the ban error",
			input: application.EventHandlerInput{
				GuildID:   shared.GuildID(1),
				UserID:    shared.UserID(2),
				ChannelID: honeypot.NONECLASS_CHANNEL_ID,
			},
			setup: func(discordService *mocks.MockDiscordService, input application.EventHandlerInput) {
				discordService.EXPECT().UserIsPrivileged(input.GuildID, input.UserID).Return(false, nil)
				discordService.EXPECT().Ban(input.GuildID, input.UserID, "Mensagem enviada em canal proibido (honeypot)", uint(1)).Return(banErr)
			},
			wantErr: banErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			discordService := mocks.NewMockDiscordService(t)
			if tt.setup != nil {
				tt.setup(discordService, tt.input)
			}

			err := honeypot.New(discordService).Execute(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Execute() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
