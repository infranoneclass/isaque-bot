package discord

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/infranoneclass/isaque-bot/internal/application"
)

type EventHandlerRegistry struct {
	state *state.State
}

func NewEventHandlerRegistry(state *state.State) *EventHandlerRegistry {
	return &EventHandlerRegistry{
		state: state,
	}
}

type EventInputMapper[T any] func(T) application.EventHandlerInput

type EventHandlerBinding interface {
	register(registry *EventHandlerRegistry)
	Intents() []gateway.Intents
}

type eventHandlerBinding[T any] struct {
	handler     application.EventHandler
	inputMapper EventInputMapper[T]
	intents     []gateway.Intents
}

func (b eventHandlerBinding[T]) register(registry *EventHandlerRegistry) {
	registry.state.AddHandler(func(data T) {
		b.handler.Execute(b.inputMapper(data))
	})
}

func (b eventHandlerBinding[T]) Intents() []gateway.Intents {
	return b.intents
}

func BindEventHandler[T any](handler application.EventHandler, inputMapper EventInputMapper[T], intents []gateway.Intents) EventHandlerBinding {
	return eventHandlerBinding[T]{
		handler:     handler,
		inputMapper: inputMapper,
		intents:     intents,
	}
}

func (r *EventHandlerRegistry) Register(bindings []EventHandlerBinding) {
	for _, binding := range bindings {
		binding.register(r)
	}
}
