package app

import (
	"context"
	"fmt"
	"go_telegram_start/internal/database"
	"go_telegram_start/internal/service/alias"
	"go_telegram_start/pkg/telegram"
	"log"
	"sync"
	"time"
)

func Run(botToken string) error {
	telegramClient := telegram.New(botToken)
	storage := database.NewLocalRedis()
	ctx := context.Background()
	offsetAsUpdateID := uint64(0)

	for {
		updates, err := telegramClient.GetUpdates(ctx, offsetAsUpdateID, 10, 0)
		if err != nil {
			return fmt.Errorf("failed at getting telegram-updates: %w", err)
		}
		if len(updates) > 0 {
			offsetAsUpdateID = updates[len(updates)-1].UpdateID + 1 // Adds 1 to get the next update
		}

		var wg sync.WaitGroup
		for _, update := range updates {
			wg.Add(1)
			go func(update telegram.Update) {
				defer wg.Done()

				game := alias.New(update, telegramClient, storage)

				if err := game.Respond(ctx); err != nil {
					log.Printf("Failed at responding to update: %+v, error: %v", update, err)
				}
			}(update)
		}
		wg.Wait()

		time.Sleep(time.Second * 5)
	}
}
