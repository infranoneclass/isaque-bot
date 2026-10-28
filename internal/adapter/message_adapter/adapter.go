package adapter

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

type MessageAdapter struct{}

func NewMessageAdapter() *MessageAdapter {
	return &MessageAdapter{}
}

func (ma *MessageAdapter) ToDiscordMessage(message string) *api.InteractionResponseData {
	return &api.InteractionResponseData{
		Content: option.NewNullableString(message),
	}
}
