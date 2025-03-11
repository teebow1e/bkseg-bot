package main

import (
	"bksecc/bkseg-bot/ctftime"
	"bksecc/bkseg-bot/discord"
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load .env")
	}

	slog.Info("starting bkseg-bot...")
	slog.Info("disgo version", slog.String("version", disgo.Version))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	ctftimeClient := ctftime.NewCTFTimeClient(5 * time.Second)

	bot := discord.NewBuddy()
	bot.Token = os.Getenv("BOT_TOKEN")
	bot.GuildID = os.Getenv("GUILD_ID")
	bot.CTFTimeClient = ctftimeClient

	if err := bot.Run(ctx); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	slog.Info("bkseg-bot is now running. Press CTRL-C to exit.")
	<-ctx.Done()

	if err := bot.Close(ctx); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
