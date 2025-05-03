package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log/slog"
)

type WordGuess struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

const rightMessage = "Верно"
const nextInWordGuessMessage = "Следующее"
const endRoundMessage = "Закончить раунд"

func NewWordGuess(tgClient *telegram.Client, u *user.User, log *slog.Logger) WordGuess {
	return WordGuess{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (w WordGuess) Respond(ctx context.Context, message string) error {
	switch message {
	case rightMessage:
		return w.saveWordResultAndGoToNextWord(ctx, user.Correct)
	case nextInWordGuessMessage:
		return w.saveWordResultAndGoToNextWord(ctx, user.Skipped)
	case endRoundMessage:
		// Check if round has ended (afterFunc has executed after round time end)
		updatedUser, err := user.NewUpdatedUser(ctx, w.user)
		if err != nil {
			return fmt.Errorf("in WordGuess failed NewUpdatedUser (userId %d): %w", w.user.TelegramID(), err)
		}
		if !updatedUser.IsStillSameGuessingRound(w.user.RoundStartTime()) { // Check if round has ended
			w.log.Info("in AfterFunc round has ended before", "user_id", w.user.TelegramID())
			_, err := w.tgClient.SendTextMessage(ctx, w.user.TelegramID(), fmt.Sprintf("Время раунда истекло"))
			if err != nil {
				return fmt.Errorf("unexpected message '%s', failed to send text message in WordGuess: %w", message, err)
			}
			return nil
		}

		roundResults, err := w.user.ConcludeRound(ctx)
		if err != nil {
			return fmt.Errorf("failed ConcludeRound for user: %d): %w", w.user.TelegramID(), err)
		}
		_, err = w.tgClient.SendTextMessage(
			ctx,
			w.user.TelegramID(),
			fmt.Sprintf("Результат раунда:\n%s", roundResults),
		)
		if err != nil {
			return fmt.Errorf("failed sending text message: %w", err)
		}
		newMenu := NewRoundResult(w.tgClient, w.user, w.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in WordGuess: %w", err)
		}
		return nil
	default:
		w.log.Debug("unknown command in WordGuess", "message", message, "user_id", w.user.TelegramID())
		_, err := w.tgClient.SendTextMessage(ctx, w.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in WordGuess: %w", message, err)
		}
		err = w.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message in WordGuess: %w", message, err)
		}
		return nil
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
	newMenu := NewWordGuess(w.tgClient, w.user, w.log)
	err = newMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sendDefaultMessage in chooseWordGuess): %w", err)
	}
	return nil
}
