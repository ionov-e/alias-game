package app

import (
	"alias-game/internal/helper/menu_factory"
	"alias-game/internal/helper/setup"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgHelper "alias-game/pkg/telegram/helper"
	tgType "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log/slog"
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
	userDB         user.DBForUserInterface
	log            *slog.Logger
	config         *setup.Config
}

func NewApp(
	tgClient *telegram.Client,
	lastUpdateIDDB lastUpdateIDDBInterface,
	userDB user.DBForUserInterface,
	log *slog.Logger,
	config *setup.Config,
) App {
	return App{
		tgClient:       tgClient,
		lastUpdateIDDB: lastUpdateIDDB,
		userDB:         userDB,
		log:            log,
		config:         config,
	}
}

func (a *App) Run(ctx context.Context) {
	offsetAsUpdateID, err := a.lastUpdateIDDB.LastUpdateID(ctx)
	if err != nil {
		a.log.Error("failed LastUpdateID", slog.String("err", err.Error()))
		return
	}

	// Буферизованный канал для задач
	taskQueue := make(chan func(), a.config.WorkerPoolSize)
	defer close(taskQueue)

	// Запуск пула воркеров
	var wg sync.WaitGroup
	for i := 0; i < a.config.WorkerPoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskQueue {
				task() // Выполнение
			}
		}()
	}

	// Основной цикл обработки обновлений
mainLoop:
	for {
		select {
		case <-ctx.Done():
			break mainLoop
		default:
			updates, err := a.tgClient.GetUpdates(ctx, offsetAsUpdateID, 10, 0)
			if err != nil {
				a.log.Error(fmt.Sprintf("failed at getting telegram-updates: %+v", err))
				continue
			}
			if len(updates) == 0 {
				time.Sleep(time.Second * 1)
				continue
			}
			offsetAsUpdateID = updates[len(updates)-1].UpdateID + 1 // Adds 1 to get the next update
			if err := a.lastUpdateIDDB.SaveLastUpdateID(ctx, offsetAsUpdateID); err != nil {
				a.log.Error(fmt.Sprintf("failed at setting lastUpdateID: %+v", err))
				continue
			}

			// Обработка каждого обновления
			for _, tgUpdate := range updates {
				select {
				case <-ctx.Done():
					break mainLoop
				default:
					task := func(tgUpdate tgType.Update) func() {
						return func() {
							tgUser, text, err := tgHelper.ExtractUserFromUpdate(tgUpdate)
							if err != nil {
								a.log.Error(
									"failed ExtractUserFromUpdate",
									slog.Any("tgUpdate", tgUpdate),
									slog.String("err", err.Error()),
								)
								return
							}
							u, err := user.NewUserFromTelegramUser(ctx, a.userDB, a.log, tgUser)
							if err != nil {
								a.log.Error("failed NewUserFromTelegramUser", slog.String("err", err.Error()))
								return
							}
							menu, err := menu_factory.MenuFactory(a.tgClient, u, a.log)
							if err != nil {
								a.log.Error("failed MenuFactory", slog.String("err", err.Error()))
								return
							}
							if err = menu.Respond(ctx, text); err != nil {
								a.log.Error(
									"failed menu.Respond",
									slog.Any("tgUpdate", tgUpdate),
									slog.String("err", err.Error()),
								)
								return
							}
						}
					}(tgUpdate)

					// Отправка задачи в канал задач
					select {
					case taskQueue <- task:
					case <-ctx.Done():
						break mainLoop
					}
				}
			}
		}
	}

	wg.Wait()

	//TODO Queue for end_round messages (results)
}
