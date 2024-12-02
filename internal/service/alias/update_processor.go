package alias

import (
	"alias-game/internal/database"
	"alias-game/internal/service/alias/menu"
	aliasUser "alias-game/internal/service/alias/user"
	"alias-game/pkg/telegram"
	"alias-game/pkg/telegram/types"
	"context"
	"fmt"
)

type UpdateProcessor struct {
	Update types.Update
	Client telegram.Client
	DB     database.DB
}

func NewUpdateProcessor(update types.Update, client telegram.Client, db database.DB) UpdateProcessor {
	return UpdateProcessor{
		Update: update,
		Client: client,
		DB:     db,
	}
}

func (up *UpdateProcessor) Respond(ctx context.Context) error {
	switch {
	case up.Update.CallbackQuery != nil:
		return up.respondToCallbackQuery(ctx, *up.Update.CallbackQuery)
	case up.Update.Message != nil:
		return up.respondToMessage(ctx, *up.Update.Message)
	default:
		return fmt.Errorf("unsupported update type: %+v", up.Update)
	}
}

func (up *UpdateProcessor) respondToMessage(ctx context.Context, message types.Message) error {
	if message.From == nil {
		return fmt.Errorf("no valid sender in message: %+v", message)
	}

	user, err := aliasUser.UserFromTelegramUser(ctx, up.DB, message.From)
	if err != nil {
		return fmt.Errorf("error getting user from Update.Message: %w", err)
	}

	ch, err := menu.FactoryMethod(user.CurrentMenuKey(), &up.Client, &user)
	if err != nil {
		return fmt.Errorf("error getting choice from Update.Message.Text: %w", err)
	}

	err = ch.Respond(ctx, message.Text)
	if err != nil {
		return fmt.Errorf("failed responding to Update.Message: %+v, error: %w", message, err)
	}
	return nil
}

func (up *UpdateProcessor) respondToCallbackQuery(ctx context.Context, callbackQuery types.CallbackQuery) error {
	user, err := aliasUser.UserFromTelegramUser(ctx, up.DB, &callbackQuery.From)
	if err != nil {
		return fmt.Errorf("error getting user from Update.CallbackQuery: %w", err)
	}

	ch, err := menu.FactoryMethod(user.CurrentMenuKey(), &up.Client, &user)
	if err != nil {
		return fmt.Errorf("error getting choice from CallbackQuery.Message.Text: %w", err)
	}

	switch m := callbackQuery.Message.(type) {
	case types.Message:
		if m.Text == "" {
			return fmt.Errorf("failed getting CallbackQuery.Message.Text: %+v", callbackQuery)
		}
		err = ch.Respond(ctx, m.Text)
		if err != nil {
			return fmt.Errorf("failed responding to CallbackQuery: %+v, error: %w", callbackQuery, err)
		}
		return nil
	case types.InaccessibleMessage:
		return fmt.Errorf("InaccessibleMessage in callbackQuery: %+v", callbackQuery)
	default:
		return fmt.Errorf("somehow no valid message in callbackQuery: %+v", callbackQuery)
	}
}
