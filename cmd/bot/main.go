package main

import (
	"context"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/joho/godotenv"
)

var commands []api.CreateCommandData = []api.CreateCommandData{
	{Name: "ping", Description: "Ping!"},
}

func main() {

	initEnv()

	r := cmdroute.NewRouter()

	r.AddFunc("ping", func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		return &api.InteractionResponseData{Content: option.NewNullableString("Pong!")}
	})

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
