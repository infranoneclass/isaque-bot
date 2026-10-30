package ping_test

import (
	"context"
	"testing"

	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/application/commands/ping"
)

func TestExecute(t *testing.T) {
	useCase := ping.New()

	response, err := useCase.Execute(context.Background(), application.CommandInput{})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if response.Content != "Pong!" {
		t.Errorf("response.Content = %q, want %q", response.Content, "Pong!")
	}

	if response.Ephemeral {
		t.Error("response.Ephemeral = true, want false")
	}
}

func TestDefinition(t *testing.T) {
	definition := ping.New().Definition()

	if definition.Name != "ping" {
		t.Errorf("definition.Name = %q, want %q", definition.Name, "ping")
	}

	if definition.Description != "Ping!" {
		t.Errorf("definition.Description = %q, want %q", definition.Description, "Ping!")
	}
}
