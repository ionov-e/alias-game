package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	userConstant "alias-game/internal/constant/user"
	userEntity "alias-game/internal/entity/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
)

type WordGuess struct {
	// Номер из словаря
	number   uint16
	tgClient *telegram.Client
	user     *userEntity.User
}

const RightMessage = "Верно"
const NextMessage = "Следующее"
const EndRoundMessage = "Закончить игру"

func NewWordGuess(wordNumber uint16, tgClient *telegram.Client, user *userEntity.User) WordGuess {
	return WordGuess{
		number:   wordNumber,
		tgClient: tgClient,
		user:     user,
	}
}

func (w WordGuess) Respond(ctx context.Context, message string) error {
	switch message {
	case RightMessage:
		return w.saveWordResultAndGoToNextWord(ctx, userConstant.Correct)
	case NextMessage:
		return w.saveWordResultAndGoToNextWord(ctx, userConstant.Skipped)
	case EndRoundMessage:
		err := chooseNewStart0(ctx, w.tgClient, w.user)
		if err != nil {
			return fmt.Errorf("failed chooseNewStart0 for user: %d): %w", w.user.TelegramID(), err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in Start0", errMessage, w.user.TelegramID())
		err := w.tgClient.SendTextMessage(ctx, w.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in Start0: %w", message, err)
		}
		err = w.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message in WordGuess: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in WordGuess", message)
	}
}

func (w WordGuess) sendDefaultMessage(ctx context.Context) error {
	word, err := w.user.Word(ctx, w.number)
	if err != nil {
		return fmt.Errorf("failed getting WordGuess (#%d): %w", w.number, err)
	}

	err = w.tgClient.SendOneTimeReplyMarkup(
		ctx,
		w.user.TelegramID(),
		word,
		tgTypes.KeyboardButtonsFromStrings(w.options()),
	)
	if err != nil {
		return fmt.Errorf("failed SendOneTimeReplyMarkup to user: %d, message %s): %w", w.user.TelegramID(), word, err)
	}

	return nil
}

func chooseWord(ctx context.Context, wordNumber uint16, client *telegram.Client, user *userEntity.User) error {
	err := user.ChangeCurrentMenu(ctx, menuConstant.NewWordKey(wordNumber))
	if err != nil {
		return fmt.Errorf("failed in chooseWord changing current thisMenu: %w", err)
	}
	thisMenu := NewWordGuess(wordNumber, client, user)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sendDefaultMessage in chooseWord): %w", err)
	}
	return nil
}

func (w WordGuess) options() []string {
	return []string{
		RightMessage,
		NextMessage,
		EndRoundMessage,
	}
}

func (w WordGuess) saveWordResultAndGoToNextWord(ctx context.Context, result userConstant.WordResult) error {
	err := w.user.UpdateWordResult(ctx, w.number, result)
	if err != nil {
		return fmt.Errorf("failed updating WordResult (#%d - %d): %w", w.number, result, err)
	}

	err = chooseWord(ctx, w.number+1, w.tgClient, w.user)
	if err != nil {
		return fmt.Errorf("failed chooseWord for user: %d): %w", w.user.TelegramID(), err)
	}
	return nil
}
