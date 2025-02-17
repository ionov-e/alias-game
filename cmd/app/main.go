package main

import (
	"alias-game/internal/app"
	"alias-game/internal/helper/setup"
	localRedis "alias-game/internal/storage/redis"
	"alias-game/pkg/telegram"
	"context"
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"log"
	"os/signal"
	"syscall" //nolint:nolintlint,goimports
)

func main() {
	if err := setup.Logging(); err != nil {
		log.Fatal(err)
	}

	tgBotToken, err := setup.TelegramBotToken()
	if err != nil {
		log.Fatal(err)
	}
	tgClient := telegram.New(tgBotToken)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default Redis
	})
	defer func() {
		err := redisClient.Close()
		if err != nil {
			log.Printf("Error closing Redis client: %v", err)
		}
	}()
	lastUpdateIDDB := localRedis.NewLastUpdateID(redisClient)
	userDB := localRedis.NewUser(redisClient)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Println("App started")

	process := app.New(tgClient, &lastUpdateIDDB, userDB)
	if err := process.Run(ctx); err != nil {
		log.Panic(err)
	}

	log.Println("App stopped")
}
