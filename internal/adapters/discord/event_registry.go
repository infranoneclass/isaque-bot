package discord

import (
	"log"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

type EventHandlerRegistry struct {
	state *state.State
}

type EventInputMapper func(*gateway.MessageCreateEvent) application.EventHandlerInput

type EventHandlerBinding struct {
	Handler     application.EventHandler
	InputMapper EventInputMapper
	Intents     []gateway.Intents
}

func NewEventHandlerRegistry(state *state.State) *EventHandlerRegistry {
	return &EventHandlerRegistry{
		state: state,
	}
}

func BindEventHandler(handler application.EventHandler, inputMapper EventInputMapper, intents []gateway.Intents) EventHandlerBinding {
	return EventHandlerBinding{
		Handler:     handler,
		InputMapper: inputMapper,
		Intents:     intents,
	}
}

func (r *EventHandlerRegistry) Register(bindings []EventHandlerBinding) {
	for _, binding := range bindings {
		r.register(binding)
	}
}

func (r *EventHandlerRegistry) register(binding EventHandlerBinding) {
	r.state.AddHandler(func(data *gateway.MessageCreateEvent) {
		log.Println("event received")
		binding.Handler.Execute(binding.InputMapper(data))
	})
}
