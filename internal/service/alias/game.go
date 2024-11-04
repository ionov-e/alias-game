package alias

import (
	"context"
	"fmt"
	"go_telegram_start/internal/database"
	"go_telegram_start/pkg/telegram"
	"go_telegram_start/pkg/telegram/types"
	"time"
)

type Game struct {
	types.Update
	telegram.Client
	database.DB
}

func New(update types.Update, client telegram.Client, db database.DB) Game {
	return Game{
		Update: update,
		Client: client,
		DB:     db,
	}
}

func (g *Game) Respond(ctx context.Context) error {
	userInfo, err := g.DB.UserInfoFromTelegramUser(ctx, g.Update.Message.User)
	if err != nil {
		return fmt.Errorf("error getting userInfo: %w", err)
	}

	newWord := "Muhaha1"

	userInfo.AddNewWord(newWord)
	userInfo.LastRequestTime = time.Now()

	err = g.DB.SaveUserInfo(ctx, userInfo)
	if err != nil {
		return fmt.Errorf("error updating userInfo: %w", err)
	}

	_, err = g.Client.SendMessage(ctx, g.Update.Message.Chat.ID, newWord)
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}

	return nil
}
