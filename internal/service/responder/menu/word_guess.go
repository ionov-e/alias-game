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
	tgClient *telegram.Client
	user     *userEntity.User
}

const rightMessage = "Верно"
const nextInWordGuessMessage = "Следующее"
const endRoundMessage = "Закончить раунд"

func NewWordGuess(tgClient *telegram.Client, user *userEntity.User) WordGuess {
	return WordGuess{
		tgClient: tgClient,
		user:     user,
	}
}

func (w WordGuess) Respond(ctx context.Context, message string) error {
	switch message {
	case rightMessage:
		return w.saveWordResultAndGoToNextWord(ctx, userConstant.Correct)
	case nextInWordGuessMessage:
		return w.saveWordResultAndGoToNextWord(ctx, userConstant.Skipped)
	case endRoundMessage:
		// TODO stop timer
		err := chooseRoundResult(ctx, w.tgClient, w.user)
		if err != nil {
			return fmt.Errorf("failed chooseRoundResult for user: %d): %w", w.user.TelegramID(), err)
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
	// TODO Send time left
	word, err := w.user.CurrentWord()
	if err != nil {
		return fmt.Errorf("failed getting CurrentWord: %w", err)
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

func chooseWordGuess(ctx context.Context, client *telegram.Client, user *userEntity.User) error {
	err := user.ChangeCurrentMenu(ctx, menuConstant.Word)
	if err != nil {
		return fmt.Errorf("failed in chooseWordGuess changing menu: %w", err)
	}
	thisMenu := NewWordGuess(client, user)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sendDefaultMessage in chooseWordGuess): %w", err)
	}
	return nil
}

func (w WordGuess) options() []string {
	return []string{
		rightMessage,
		nextInWordGuessMessage,
		endRoundMessage,
	}
}

func (w WordGuess) saveWordResultAndGoToNextWord(ctx context.Context, result userConstant.WordResult) error {
	w.user.SetCurrentWordResult(result)
	w.user.NextWord()
	err := chooseWordGuess(ctx, w.tgClient, w.user)
	if err != nil {
		return fmt.Errorf("failed chooseWordGuess for user: %d): %w", w.user.TelegramID(), err)
	}
	return nil
}
