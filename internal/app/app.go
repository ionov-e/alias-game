package app

import (
	userEntity "alias-game/internal/entity/user"
	"alias-game/internal/helper"
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

func (a *App) Run(ctx context.Context) error { // TODO no return
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
	loopUpdates:
		for _, tgUpdate := range updates {
			select {
			case <-ctx.Done():
				break loopUpdates
			default:
				wg.Add(1)
				go func() {
					defer wg.Done()

					tgUser, text, err := helper.ExtractFromUpdate(tgUpdate)
					if err != nil {
						log.Printf("failed at extracting from tgUpdate: %+v, error: %v", tgUpdate, err)
						return
					}

					user, err := userEntity.NewFromTelegramUser(ctx, a.userDB, tgUser)
					if err != nil {
						log.Printf("error getting user from Update.CallbackQuery: %v", err)
						return
					}

					menu, err := helper.MenuFactory(a.tgClient, user)
					if err != nil {
						log.Printf("error getting choice from CallbackQuery.Message.Text: %v", err)
						return
					}

					if err = menu.Respond(ctx, text); err != nil {
						log.Printf("failed at responding to tgUpdate: %+v, error: %v", tgUpdate, err)
						return
					}
				}()
			}
		}
		wg.Wait() //TODO think about limit (worker)
	}

	//TODO Queue for end_round messages (results)
}
