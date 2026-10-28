package main

import (
	"context"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	adapter "github.com/infranoneclass/isaque-bot/internal/adapter/message_adapter"
	pingcommand "github.com/infranoneclass/isaque-bot/internal/command/ping"
	pinghandler "github.com/infranoneclass/isaque-bot/internal/handler/ping"

	"github.com/joho/godotenv"
)

var commands []api.CreateCommandData = []api.CreateCommandData{
	{Name: "ping", Description: "Ping!"},
}

func main() {

	initEnv()

	r := cmdroute.NewRouter()
	registerCommands(r)

	s := state.New("Bot " + os.Getenv("BOT_PRIVATE_TOKEN"))

	s.AddInteractionHandler(r)
	s.AddIntents(gateway.IntentGuilds)

	if err := cmdroute.OverwriteCommands(s, commands); err != nil {
		log.Fatalln("cannot update commands:", err)
	}

	log.Println("Connecting bot")
	if err := s.Connect(context.TODO()); err != nil {
		log.Println("cannot connect:", err)
	}
}

func initEnv() {
	botPrivateToken := os.Getenv("BOT_PRIVATE_TOKEN")
	if botPrivateToken == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func registerCommands(r *cmdroute.Router) {
	messageAdapter := adapter.NewMessageAdapter()
	pingHandler := pinghandler.NewPingHandler(
		messageAdapter,
		pingcommand.NewPingCommand(),
	)

	r.AddFunc("ping", pingHandler.HandleCommand)
}
