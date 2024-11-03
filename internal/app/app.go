package app

import (
	"context"
	"fmt"
	"go_telegram_start/internal/database"
	"go_telegram_start/internal/service/alias"
	"go_telegram_start/pkg/telegram"
	"time"
)

func Run(botToken string) error {
	telegramClient := telegram.New(botToken)
	storage := database.NewLocalRedis()
	ctx := context.Background()
	lastUpdateID := uint64(0)

	for {
		updates, err := telegramClient.GetUpdates(ctx, lastUpdateID, 10, 0)
		if err != nil {
			return fmt.Errorf("failed at getting telegram-updates: %w", err)
		}
		for _, update := range updates {
			if update.UpdateID == lastUpdateID {
				continue
			}
			lastUpdateID = update.UpdateID
			game := alias.New(update, telegramClient, storage)
			if err := game.Respond(ctx); err != nil {
				return fmt.Errorf("failde at responding to update: %w", err)
			}
		}
		time.Sleep(time.Second * 5)
	}
}
