package bootstrap

import (
	"github.com/infranoneclass/isaque-bot/internal/adapters/discord"
	"github.com/infranoneclass/isaque-bot/internal/repositories"
)

func NewDiscordBot(configs *repositories.Config) *discord.Bot {
	return discord.NewBot(configs)
}
