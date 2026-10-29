package bootstrap

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
}

func LoadConfig() (Config, error) {
	if strings.TrimSpace(os.Getenv("BOT_PRIVATE_TOKEN")) == "" {
		if err := godotenv.Load(); err != nil {
			return Config{}, fmt.Errorf("load environment file: %w", err)
		}
	}

	config := Config{BotToken: strings.TrimSpace(os.Getenv("BOT_PRIVATE_TOKEN"))}
	if config.BotToken == "" {
		return Config{}, fmt.Errorf("BOT_PRIVATE_TOKEN is required")
	}

	return config, nil
}
