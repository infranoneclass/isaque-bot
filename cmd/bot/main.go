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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signals)

	go func() {
		sig := <-signals
		log.Printf("Received %s; shutting down bot", sig)
		cancel()
	}()

	log.Println("Connecting bot")
	if err := bot.Run(ctx); err != nil {
		log.Printf("Discord bot stopped with an error: %v", err)
		return
	}

	if ctx.Err() == nil {
		log.Println("Discord bot stopped unexpectedly without an error")
		return
	}

	log.Println("Bot stopped cleanly")
}
