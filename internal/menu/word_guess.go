package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
)

type WordGuess struct {
	tgClient *telegram.Client
	user     *user.User
}

const rightMessage = "Верно"
const nextInWordGuessMessage = "Следующее"
const endRoundMessage = "Закончить раунд"

func NewWordGuess(tgClient *telegram.Client, u *user.User) WordGuess {
	return WordGuess{
		tgClient: tgClient,
		user:     u,
	}
}

func (w WordGuess) Respond(ctx context.Context, message string) error {
	switch message {
	case rightMessage:
		return w.saveWordResultAndGoToNextWord(ctx, user.Correct)
	case nextInWordGuessMessage:
		return w.saveWordResultAndGoToNextWord(ctx, user.Skipped)
	case endRoundMessage:
		// TODO stop timer
		err := w.user.ChangeCurrentMenu(ctx, menuConstant.RoundResult)
		if err != nil {
			return fmt.Errorf("failed in WordGuess changing current menu: %w", err)
		}
		roundResults, err := w.user.ConcludeRound(ctx)
		if err != nil {
			return fmt.Errorf("failed ConcludeRound for user: %d): %w", w.user.TelegramID(), err)
		}
		err = w.tgClient.SendTextMessage(
			ctx,
			w.user.TelegramID(),
			fmt.Sprintf("Результат раунда:\n%s", roundResults),
		)
		if err != nil {
			return fmt.Errorf("failed sending text message: %w", err)
		}
		newMenu := NewRoundResult(w.tgClient, w.user)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in WordGuess: %w", err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in WordGuess", errMessage, w.user.TelegramID())
		err := w.tgClient.SendTextMessage(ctx, w.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in WordGuess: %w", message, err)
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
		tgTypes.KeyboardButtonsFromStrings([]string{rightMessage, nextInWordGuessMessage, endRoundMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed in WordGuess SendOneTimeReplyMarkup to user: %d, message %s): %w", w.user.TelegramID(), word, err)
	}

	return nil
}

func (w WordGuess) saveWordResultAndGoToNextWord(ctx context.Context, result user.WordResult) error {
	w.user.SetCurrentWordResult(result)
	w.user.NextWord()
	err := w.user.ChangeCurrentMenu(ctx, menuConstant.Word)
	if err != nil {
		return fmt.Errorf("failed in chooseWordGuess changing menu: %w", err)
	}
	newMenu := NewWordGuess(w.tgClient, w.user)
	err = newMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sendDefaultMessage in chooseWordGuess): %w", err)
	}
	return nil
}
