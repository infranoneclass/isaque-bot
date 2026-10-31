package honeypot_test

import (
	"testing"

	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/application/events/honeypot"
)

func TestExecute(t *testing.T) {
	useCase := honeypot.New()

	err := useCase.Execute(application.EventHandlerInput{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
}
