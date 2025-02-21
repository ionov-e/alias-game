package main

import (
	"alias-game/internal/app"
	"alias-game/internal/helper/setup"
	"alias-game/internal/last_update_id"
	"alias-game/internal/user"
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
	//TODO config
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
	dbForLastUpdateID := last_update_id.NewLastUpdateIDRedisClient(redisClient)
	dbForUser := user.NewRedisClient(redisClient)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Println("App started")

	process := app.New(tgClient, &dbForLastUpdateID, dbForUser)
	if err := process.Run(ctx); err != nil {
		log.Panic(err)
	}

	log.Println("App stopped")
}
