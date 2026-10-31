package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/infranoneclass/isaque-bot/internal/bootstrap"
)

func main() {

	config, err := bootstrap.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot := bootstrap.NewDiscordBot(config)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer func() {
		log.Println("Kill signal received. Shutting bot down")
		stop()
	}()

	log.Println("Connecting bot")
	if err := bot.Run(ctx); err != nil {
		log.Println("Cannot connect:", err)
	}
}
