package alias

import (
	"alias-game/internal/database"
	"alias-game/pkg/telegram"
	"alias-game/pkg/telegram/types"
	"context"
	"fmt"
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
	user, err := UserFromTelegramUser(ctx, g.DB, *g.Update.Message.From)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	newWord := "Muhaha1"

	if err = user.AddNewWord(ctx, newWord); err != nil {
		return fmt.Errorf("error adding new word into user: %w", err)
	}

	_, err = g.Client.SendMessage(ctx, types.SendMessage{ChatID: g.Update.Message.Chat.ID, Text: newWord})
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}

	return nil
}
