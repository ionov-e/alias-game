package app

import (
	"alias-game/internal/database"
	"alias-game/internal/service/alias"
	"alias-game/pkg/telegram"
	"alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func Run(botToken string) error {
	telegramClient := telegram.New(botToken)
	storage := database.NewLocalRedis()
	ctx := context.Background()
	offsetAsUpdateID, err := storage.LastUpdateID(ctx)
	if err != nil {
		return fmt.Errorf("failed at getting lastUpdateID: %w", err)
	}

	for {
		updates, err := telegramClient.GetUpdates(ctx, offsetAsUpdateID, 10, 0)
		if err != nil {
			return fmt.Errorf("failed at getting telegram-updates: %w", err)
		}
		if len(updates) > 0 {
			offsetAsUpdateID = updates[len(updates)-1].UpdateID + 1 // Adds 1 to get the next update

			if err := storage.SaveLastUpdateID(ctx, offsetAsUpdateID); err != nil {
				return fmt.Errorf("failed at setting lastUpdateID: %w", err)
			}
		}

		var wg sync.WaitGroup
		for _, update := range updates {
			wg.Add(1)
			go func(update types.Update) {
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
