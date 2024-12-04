package main

import (
	"alias-game/internal/app"
	"alias-game/internal/database"
	"alias-game/internal/setup"
	"alias-game/pkg/telegram"
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	if err := setup.Logging(); err != nil {
		log.Fatal(err)
	}

	botToken, err := setup.Token()
	if err != nil {
		log.Fatal(err)
	}

	telegramClient := telegram.New(botToken)

	storage := database.NewLocalRedis()
	defer func() {
		if redisClient, ok := storage.(*database.Redis); ok {
			if err := redisClient.Close(); err != nil {
				log.Printf("Error closing Redis client: %v", err)
			}
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Println("App started")

	if err := app.Run(ctx, telegramClient, storage); err != nil {
		log.Panic(err)
	}

	log.Println("App stopped")
}
