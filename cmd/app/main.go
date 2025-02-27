package main

import (
	"alias-game/internal/app"
	"alias-game/internal/helper/setup"
	"alias-game/internal/last_update_id"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	"context"
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"os/signal"
	"syscall" //nolint:nolintlint,goimports
)

func main() {
	logger := setup.GetLogger(true)
	//TODO config
	tgBotToken, err := setup.TelegramBotToken(logger)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	tgClient := telegram.NewClient(tgBotToken, logger)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default Redis
	})
	defer func() {
		err := redisClient.Close()
		if err != nil {
			logger.Error("Error closing Redis client", "err", err)
		}
	}()
	dbForLastUpdateID := last_update_id.NewLastUpdateIDRedisClient(redisClient, logger)
	dbForUser := user.NewRedisClient(redisClient, logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Info("App started")

	process := app.NewApp(tgClient, &dbForLastUpdateID, dbForUser, logger)
	process.Run(ctx)

	logger.Info("App stopped")
}
