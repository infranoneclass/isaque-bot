package discord

import (
	"github.com/diamondburned/arikawa/v3/api"
	discordmodel "github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

type ResponsePresenter struct{}

func NewResponsePresenter() *ResponsePresenter {
	return &ResponsePresenter{}
}

func (p *ResponsePresenter) ToDiscord(response application.Response) *api.InteractionResponseData {
	return &api.InteractionResponseData{
		Content: option.NewNullableString(response.Content),
		Flags:   responseFlags(response),
	}
}

func (p *ResponsePresenter) Error() *api.InteractionResponseData {
	return p.ToDiscord(application.Response{
		Content:   "Algo deu errado!",
		Ephemeral: true,
	})
}

func responseFlags(response application.Response) discordmodel.MessageFlags {
	if response.Ephemeral {
		return discordmodel.EphemeralMessage
	}

	return 0
}
