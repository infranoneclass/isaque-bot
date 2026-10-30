package discord

import "github.com/diamondburned/arikawa/v3/gateway"

func MergeIntents(intents ...gateway.Intents) gateway.Intents {
	var merged gateway.Intents

	for _, intent := range intents {
		merged |= intent
	}

	return merged
}
