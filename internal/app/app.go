package app

import (
	"alias-game/internal/service/responder"
	"alias-game/internal/storage"
	"alias-game/pkg/telegram"
	"context"
	"fmt" //nolint:goimports
	"log"
	"sync"
	"time"
)

type lastUpdateIDDBInterface interface {
	LastUpdateID(ctx context.Context) (uint64, error)
	SaveLastUpdateID(ctx context.Context, lastUpdateID uint64) error
}

type App struct {
	tgClient       *telegram.Client
	lastUpdateIDDB lastUpdateIDDBInterface
	userDB         storage.UserDBInterface
}

func New(
	tgClient *telegram.Client,
	lastUpdateIDDB lastUpdateIDDBInterface,
	userDB storage.UserDBInterface,
) App {
	return App{
		tgClient:       tgClient,
		lastUpdateIDDB: lastUpdateIDDB,
		userDB:         userDB,
	}
}

func (a *App) Run(ctx context.Context) error {
	offsetAsUpdateID, err := a.lastUpdateIDDB.LastUpdateID(ctx)
	if err != nil {
		return fmt.Errorf("failed at getting lastUpdateID: %w", err)
	}

	for {
		updates, err := a.tgClient.GetUpdates(ctx, offsetAsUpdateID, 10, 0)
		if err != nil {
			return fmt.Errorf("failed at getting telegram-updates: %w", err)
		}
		if len(updates) == 0 {
			time.Sleep(time.Second * 1)
			continue
		}

		offsetAsUpdateID = updates[len(updates)-1].UpdateID + 1 // Adds 1 to get the next update

		if err := a.lastUpdateIDDB.SaveLastUpdateID(ctx, offsetAsUpdateID); err != nil {
			return fmt.Errorf("failed at setting lastUpdateID: %w", err)
		}

		var wg sync.WaitGroup
		for _, updateResponse := range updates {
			wg.Add(1)
			go func() {
				defer wg.Done()

				process := responder.New(updateResponse, a.tgClient, a.userDB)
				if err = process.Run(ctx); err != nil {
					log.Printf("Failed at responding to update: %+v, error: %v", updateResponse, err)
				}
			}()
		}
		wg.Wait() //TODO think about limit (worker)
	}

	//TODO Queue for end_round messages (results)
}
