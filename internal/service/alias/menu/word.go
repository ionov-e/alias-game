package menu

import (
	dbConstants "alias-game/internal/database/constants"
	aliasUser "alias-game/internal/service/alias/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
)

type Word struct {
	// Номер из словаря
	number uint16
	client *telegram.Client
	user   *aliasUser.User
}

const RightMessage = "Верно"
const NextMessage = "Следующее"
const EndRoundMessage = "Закончить игру"

func NewWord(wordNumber uint16, client *telegram.Client, user *aliasUser.User) Menu {
	return Word{
		number: wordNumber,
		client: client,
		user:   user,
	}
}

func ChooseWord(ctx context.Context, wordNumber uint16, client *telegram.Client, user *aliasUser.User) error {
	thisMenu := NewWord(wordNumber, client, user)
	err := user.ChangeCurrentMenu(ctx, thisMenu.RedisKey())
	if err != nil {
		return fmt.Errorf("failed in ChooseWord changing current thisMenu: %w", err)
	}
	err = thisMenu.DefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed DefaultMessage in ChooseWord): %w", err)
	}
	return nil
}

func (w Word) DefaultMessage(ctx context.Context) error {
	word, err := w.user.Word(ctx, w.number)
	if err != nil {
		return fmt.Errorf("failed getting Word (#%d - %d): %w", w.number, dbConstants.Correct, err)
	}

	err = w.client.SendOneTimeReplyMarkup(
		ctx,
		w.user.TelegramID(),
		word,
		tgTypes.KeyboardButtonsFromStrings(w.Options()),
	)
	if err != nil {
		return fmt.Errorf("failed SendOneTimeReplyMarkup to user: %d, message %s): %w", w.user.TelegramID(), word, err)
	}

	return nil
}

func (w Word) RedisKey() dbConstants.MenuKeyStored {
	return dbConstants.NewWordMenuKey(w.number)
}

func (w Word) Options() []string {
	return []string{
		RightMessage,
		NextMessage,
		EndRoundMessage,
	}
}

func (w Word) Respond(ctx context.Context, message string) error {
	switch message {
	case RightMessage:
		err := w.user.UpdateWordResult(ctx, w.number, dbConstants.Correct)
		if err != nil {
			return fmt.Errorf("failed updating WordResult (#%d - %d): %w", w.number, dbConstants.Correct, err)
		}

		err = ChooseWord(ctx, w.number+1, w.client, w.user)
		if err != nil {
			return fmt.Errorf("failed ChooseWord for user: %d): %w", w.user.TelegramID(), err)
		}
		return nil
	case NextMessage:
		err := w.user.UpdateWordResult(ctx, w.number, dbConstants.Skipped)
		if err != nil {
			return fmt.Errorf("failed updating WordResult (#%d - %d): %w", w.number, dbConstants.Skipped, err)
		}

		err = ChooseWord(ctx, w.number+1, w.client, w.user)
		if err != nil {
			return fmt.Errorf("failed ChooseWord for user: %d): %w", w.user.TelegramID(), err)
		}
		return nil
	case EndRoundMessage:
		err := ChooseNewStart0(ctx, w.client, w.user)
		if err != nil {
			return fmt.Errorf("failed ChooseNewStart0 for user: %d): %w", w.user.TelegramID(), err)
		}
		return nil
	default:
		err := w.DefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message in Word: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in Word", message)
	}
}
