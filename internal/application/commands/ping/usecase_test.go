package ping_test

import (
	"context"
	"testing"

	"github.com/infranoneclass/isaque-bot/internal/application"
	"github.com/infranoneclass/isaque-bot/internal/application/commands/ping"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name  string
		input application.CommandInput
		want  application.Response
	}{
		{
			name:  "without options",
			input: application.CommandInput{},
			want:  application.Response{Content: "Pong!"},
		},
		{
			name:  "with options",
			input: application.CommandInput{Options: map[string]string{"ignored": "value"}},
			want:  application.Response{Content: "Pong!"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := ping.New().Execute(context.Background(), tt.input)
			if err != nil {
				t.Fatalf("Execute() error = %v", err)
			}

			if response != tt.want {
				t.Errorf("Execute() response = %+v, want %+v", response, tt.want)
			}
		})
	}
}

func TestDefinition(t *testing.T) {
	tests := []struct {
		name string
		want application.CommandDefinition
	}{
		{
			name: "ping command definition",
			want: application.CommandDefinition{
				Name:        "ping",
				Description: "Ping!",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if definition := ping.New().Definition(); definition != tt.want {
				t.Errorf("Definition() = %+v, want %+v", definition, tt.want)
			}
		})
	}
}
