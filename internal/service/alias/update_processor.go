package alias

import (
	"alias-game/internal/database"
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
	user, err := UserFromTelegramUser(ctx, up.DB, message.From)
	if err != nil {
		return fmt.Errorf("error getting user from Update.Message: %w", err)
	}
	err = up.sendOneTimeReplyMarkup(ctx, user, "haha-Message", [][]types.KeyboardButton{{{Text: "Option 1"}, {Text: "Option 2"}}})
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	return nil
}

func (up *UpdateProcessor) respondToCallbackQuery(ctx context.Context, callbackQuery types.CallbackQuery) error {
	user, err := UserFromTelegramUser(ctx, up.DB, &callbackQuery.From)
	if err != nil {
		return fmt.Errorf("error getting user from Update.CallbackQuery: %w", err)
	}
	err = up.sendOneTimeReplyMarkup(ctx, user, "haha-CallbackQuery", [][]types.KeyboardButton{{{Text: "Option 1"}, {Text: "Option 2"}}})
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	return nil
}

func (up *UpdateProcessor) sendOneTimeReplyMarkup(ctx context.Context, user User, text string, keyboardButtons [][]types.KeyboardButton) error {
	_, err := up.Client.SendMessage(ctx, types.SendMessage{
		ChatID: user.userInfo.TelegramID,
		Text:   text,
		ReplyMarkup: types.ReplyKeyboardMarkup{
			OneTimeKeyboard: true,
			Keyboard:        keyboardButtons,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send reply markup: user=%+v, text=%s, keyboardButtons=%+v, error=%w", user.userInfo.TelegramID, text, keyboardButtons, err)
	}
	return nil
}
