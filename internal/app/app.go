package app

import (
	"alias-game/internal/database"
	"alias-game/internal/service/alias"
	"alias-game/pkg/telegram"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func Run(ctx context.Context, telegramClient telegram.Client, storage database.DB) error {
	offsetAsUpdateID, err := storage.LastUpdateID(ctx)
	if err != nil {
		return fmt.Errorf("failed at getting lastUpdateID: %w", err)
	}

	for {
		updates, err := telegramClient.GetUpdates(ctx, offsetAsUpdateID, 10, 0)
		if err != nil {
			return fmt.Errorf("failed at getting telegram-updates: %w", err)
		}
		if len(updates) == 0 {
			time.Sleep(time.Second * 1)
			continue
		}

		offsetAsUpdateID = updates[len(updates)-1].UpdateID + 1 // Adds 1 to get the next update

		if err := storage.SaveLastUpdateID(ctx, offsetAsUpdateID); err != nil {
			return fmt.Errorf("failed at setting lastUpdateID: %w", err)
		}

		var wg sync.WaitGroup
		for _, update := range updates {
			wg.Add(1)
			go func() {
				defer wg.Done()

				updateProcessor := alias.NewUpdateProcessor(update, telegramClient, storage)

				if err := updateProcessor.Respond(ctx); err != nil {
					log.Printf("Failed at responding to update: %+v, error: %v", update, err)
				}
			}()
		}
		wg.Wait() //TODO think about limit (worker)
	}

	//TODO Queue for end_round messages (results)
}
