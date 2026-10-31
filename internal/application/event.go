package application

type EventHandler interface {
	Execute(input EventHandlerInput)
}

type EventHandlerInput struct {
	Options map[string]any
}
