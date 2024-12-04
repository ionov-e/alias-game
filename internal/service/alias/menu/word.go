package menu

import (
	dbConstants "alias-game/internal/database/constants"
	aliasUser "alias-game/internal/service/alias/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
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
	err := user.ChangeCurrentMenu(ctx, dbConstants.NewWordMenuKey(wordNumber))
	if err != nil {
		return fmt.Errorf("failed in ChooseWord changing current thisMenu: %w", err)
	}
	thisMenu := NewWord(wordNumber, client, user)
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
		return w.saveWordResultAndGoToNextWord(ctx, dbConstants.Correct)
	case NextMessage:
		return w.saveWordResultAndGoToNextWord(ctx, dbConstants.Skipped)
	case EndRoundMessage:
		err := ChooseNewStart0(ctx, w.client, w.user)
		if err != nil {
			return fmt.Errorf("failed ChooseNewStart0 for user: %d): %w", w.user.TelegramID(), err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in Start0", errMessage, w.user.TelegramID())
		err := w.client.SendTextMessage(ctx, w.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in Start0: %w", message, err)
		}
		err = w.DefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message in Word: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in Word", message)
	}
}

func (w Word) saveWordResultAndGoToNextWord(ctx context.Context, result dbConstants.WordResult) error {
	err := w.user.UpdateWordResult(ctx, w.number, result)
	if err != nil {
		return fmt.Errorf("failed updating WordResult (#%d - %d): %w", w.number, result, err)
	}

	err = ChooseWord(ctx, w.number+1, w.client, w.user)
	if err != nil {
		return fmt.Errorf("failed ChooseWord for user: %d): %w", w.user.TelegramID(), err)
	}
	return nil
}
