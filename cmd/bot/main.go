package main

import (
	"context"
	"log"

	"github.com/infranoneclass/isaque-bot/internal/bootstrap"
)

func main() {

	config, err := bootstrap.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot := bootstrap.NewDiscordBot(config)

	log.Println("Connecting bot")
	if err := bot.Run(context.Background()); err != nil {
		log.Println("cannot connect:", err)
	}
}
