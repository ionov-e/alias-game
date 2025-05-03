package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type WordGuess struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

const rightMessage = "Верно"
const nextInWordGuessMessage = "Следующее"
const endRoundMessage = "Закончить раунд"

const timeLeftInSecondsMessage = "⏱️Осталось секунд"

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
	messageWithTimeLeft, err := w.tgClient.SendTextMessage(ctx, w.user.TelegramID(), fmt.Sprintf("%s: %d", timeLeftInSecondsMessage, w.user.RoundTimeLeftInSeconds()))
	if err != nil {
		return fmt.Errorf("failed to send text messageWithTimeLeft in WordGuess: %w", err)
	}
	w.updateOrDeleteEveryFewSecondsMessageWithTimeLeft(ctx, w.user.CurrenRoundWordNumber(), messageWithTimeLeft.Result.MessageID)

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

func (w WordGuess) updateOrDeleteEveryFewSecondsMessageWithTimeLeft(
	ctx context.Context,
	roundWordNumberForMessage uint16,
	messageIDWithTimeLeft int64,
) {
	time.AfterFunc(3*time.Second, func() {
		updatedUser, err := user.NewUpdatedUser(ctx, w.user)
		if err != nil {
			w.log.Error("in updateOrDeleteEveryFewSecondsMessageWithTimeLeft failed NewUpdatedUser: %w", err, "user_id", w.user.TelegramID())
			return
		}

		// Delete message if round has ended or new word has been already given
		if !updatedUser.IsStillSameGuessingRound(w.user.RoundStartTime()) || roundWordNumberForMessage != updatedUser.CurrenRoundWordNumber() {
			_, err = w.tgClient.DeleteMessage(ctx, w.user.TelegramID(), messageIDWithTimeLeft)
			if err != nil {
				w.log.Error("failed to delete messageWithTimeLeft in WordGuess", "error", err)
			}
			return
		}

		// Update message
		_, err = w.tgClient.UpdateMessageText(ctx, w.user.TelegramID(), messageIDWithTimeLeft, fmt.Sprintf("%s: %d", timeLeftInSecondsMessage, updatedUser.RoundTimeLeftInSeconds()))
		if err != nil {
			w.log.Error("failed to delete messageWithTimeLeft in WordGuess", "error", err)
		}
		w.updateOrDeleteEveryFewSecondsMessageWithTimeLeft(ctx, roundWordNumberForMessage, messageIDWithTimeLeft)
	})
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
