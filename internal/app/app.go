package app

import (
	"alias-game/internal/helper/menu_factory"
	"alias-game/internal/helper/setup"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgHelper "alias-game/pkg/telegram/helper"
	tgType "alias-game/pkg/telegram/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

const webhookUrlPath = "/webhook"

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
	var wg sync.WaitGroup
	concurrencyLimiter := make(chan struct{}, a.config.ConcurrencyLimit)
	defer close(concurrencyLimiter)

	if !a.config.Webhook.Enabled {
		a.runPolling(ctx, &wg, concurrencyLimiter)
		return
	}

	a.enableWebhook(ctx, &wg, concurrencyLimiter)
}

func (a *App) runPolling(ctx context.Context, wg *sync.WaitGroup, concurrencyLimiter chan struct{}) {
	offsetAsUpdateID, err := a.lastUpdateIDDB.LastUpdateID(ctx)
	if err != nil {
		a.log.Error("failed LastUpdateID", slog.String("err", err.Error()))
		return
	}
	err = a.tgClient.DeleteWebhook(ctx)
	if err != nil {
		a.log.Error("failed DeleteWebhook", slog.String("err", err.Error()))
		return
	}

	if !a.checkWithTestMessage(ctx, "polling") {
		return
	}

mainLoop:
	for {
		select {
		case <-ctx.Done():
			break mainLoop
		default:
			updates, err := a.tgClient.GetUpdates(ctx, offsetAsUpdateID, 10, 0)
			if err != nil {
				a.log.Error("failed at getting telegram-updates", slog.String("err", err.Error()))
				continue
			}
			if len(updates) == 0 {
				time.Sleep(time.Second)
				continue
			}
			offsetAsUpdateID = updates[len(updates)-1].UpdateID + 1
			if err := a.lastUpdateIDDB.SaveLastUpdateID(ctx, offsetAsUpdateID); err != nil {
				a.log.Error("failed to save lastUpdateID", slog.String("err", err.Error()))
				continue
			}
			for _, tgUpdate := range updates {
				select {
				case <-ctx.Done():
					break mainLoop
				case concurrencyLimiter <- struct{}{}:
					wg.Add(1)
					go a.processUpdate(ctx, tgUpdate, wg, concurrencyLimiter)
				}
			}
		}
	}

	wg.Wait()
}

func (a *App) enableWebhook(ctx context.Context, wg *sync.WaitGroup, concurrencyLimiter chan struct{}) {
	webhookURL := fmt.Sprintf("%s/webhook", a.config.Webhook.URL)
	a.log.Debug("setting webhook", slog.String("url", webhookURL))
	err := a.tgClient.SetWebhook(ctx, webhookURL)
	if err != nil {
		a.log.Error("failed to set webhook", slog.String("err", err.Error()))
		return
	}

	go a.runWebhook(ctx, wg, concurrencyLimiter)

	if !a.checkWithTestMessage(ctx, "webhook") {
		return
	}

	<-ctx.Done()
	wg.Wait()
}

func (a *App) runWebhook(ctx context.Context, wg *sync.WaitGroup, concurrencyLimiter chan struct{}) {
	mux := http.NewServeMux()
	mux.HandleFunc(webhookUrlPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var tgUpdate tgType.Update
		if err := json.NewDecoder(r.Body).Decode(&tgUpdate); err != nil {
			a.log.Error("failed to decode webhook update", slog.String("err", err.Error()))
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		select {
		case concurrencyLimiter <- struct{}{}:
			wg.Add(1)
			go a.processUpdate(ctx, tgUpdate, wg, concurrencyLimiter)
			w.WriteHeader(http.StatusOK)
		default:
			a.log.Warn("webhook update dropped due to concurrency limit")
			http.Error(w, "too many requests", http.StatusTooManyRequests)
		}
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.config.Webhook.Port),
		Handler: mux,
	}

	go func() {
		a.log.Info("starting webhook server", slog.String("port", a.config.Webhook.Port))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Error("webhook server error", slog.String("err", err.Error()))
		}
	}()

	go func() {
		<-ctx.Done()
		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctxTimeout); err != nil {
			a.log.Error("failed to shutdown webhook server", slog.String("err", err.Error()))
		}
	}()
}

func (a *App) checkWithTestMessage(ctx context.Context, method string) bool {
	testChatID := a.config.Telegram.TestChatID
	if testChatID == 0 {
		a.log.Warn(method + " NOT verified: no test chat id provided")
		return true
	}

	if _, err := a.tgClient.SendTextMessage(ctx, testChatID, "✅ "+method); err != nil {
		a.log.Error(method+" test message failed", slog.String("err", err.Error()))
		return false
	}
	a.log.Info(method+" verified via test message", slog.Int64("chat_id", testChatID))
	return true
}

func (a *App) processUpdate(ctx context.Context, tgUpdate tgType.Update, wg *sync.WaitGroup, limiter chan struct{}) {
	defer func() {
		wg.Done()
		<-limiter
	}()

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
