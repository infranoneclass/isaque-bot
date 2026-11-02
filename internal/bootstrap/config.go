package bootstrap

import (
	"fmt"
	"os"
	"strings"

	"github.com/infranoneclass/isaque-bot/internal/repositories"
	"github.com/joho/godotenv"
)

func LoadConfig() (*repositories.Config, error) {
	config := repositories.NewConfigRepository()

	if strings.TrimSpace(os.Getenv("BOT_PRIVATE_TOKEN")) == "" {
		if err := godotenv.Load(); err != nil {
			return config, fmt.Errorf("load environment file: %w", err)
		}
	}

	config.BindBotToken(strings.TrimSpace(os.Getenv("BOT_PRIVATE_TOKEN")))
	if config.BotToken == "" {
		return config, fmt.Errorf("BOT_PRIVATE_TOKEN is required")
	}

	jsonConfigFile, err := os.ReadFile("configs.json")
	if err != nil {
		return config, fmt.Errorf("could not load configs file: %e", err)
	}
	err = config.BindFromJSONFile(jsonConfigFile)

	return config, nil
}
