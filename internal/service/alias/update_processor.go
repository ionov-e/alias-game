package alias

import (
	"alias-game/internal/database"
	"alias-game/internal/service/alias/menu"
	aliasUser "alias-game/internal/service/alias/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
)

type UpdateProcessor struct {
	Update tgTypes.Update
	Client telegram.Client
	DB     database.DB
}

func NewUpdateProcessor(update tgTypes.Update, client telegram.Client, db database.DB) UpdateProcessor {
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

func (up *UpdateProcessor) respondToMessage(ctx context.Context, message tgTypes.Message) error {
	if message.From == nil {
		return fmt.Errorf("no user in Message.From: %+v", message)
	}
	err := up.respondToText(ctx, message.From, message.Text)
	if err != nil {
		return fmt.Errorf("failed responding to Message: %+v, error: %w", message, err)
	}

	return nil
}

func (up *UpdateProcessor) respondToCallbackQuery(ctx context.Context, callbackQuery tgTypes.CallbackQuery) error {
	switch message := callbackQuery.Message.(type) {
	case tgTypes.Message:
		if message.Text == "" {
			return fmt.Errorf("failed getting CallbackQuery.Message.Text: %+v", callbackQuery)
		}
		if &callbackQuery.From == nil {
			return fmt.Errorf("no user (impossible by contract) in CallbackQuery.From: %+v", callbackQuery)
		}
		err := up.respondToText(ctx, &callbackQuery.From, message.Text)
		if err != nil {
			return fmt.Errorf("failed responding to CallbackQuery: %+v, error: %w", callbackQuery, err)
		}
		return nil
	case tgTypes.InaccessibleMessage:
		return fmt.Errorf("InaccessibleMessage in callbackQuery: %+v", callbackQuery)
	default:
		return fmt.Errorf("somehow no valid message in callbackQuery: %+v", callbackQuery)
	}
}

func (up *UpdateProcessor) respondToText(ctx context.Context, userFromTelegram *tgTypes.User, text string) error {
	user, err := aliasUser.UserFromTelegramUser(ctx, up.DB, userFromTelegram)
	if err != nil {
		return fmt.Errorf("error getting user from Update.CallbackQuery: %w", err)
	}

	currentMenu, err := menu.FactoryMethod(user.CurrentMenuKey(), &up.Client, &user)
	if err != nil {
		return fmt.Errorf("error getting choice from CallbackQuery.Message.Text: %w", err)
	}

	return currentMenu.Respond(ctx, text) //nolint:wrapcheck
}
