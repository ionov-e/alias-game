package menu

import (
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log/slog"
	"time"
)

const startNewRoundMessage = "Начать раунд"

type NextRoundSuggestion struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

func NewNextRoundSuggestion(tgClient *telegram.Client, u *user.User, log *slog.Logger) NextRoundSuggestion {
	return NextRoundSuggestion{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (m NextRoundSuggestion) Respond(ctx context.Context, message string) error {
	switch message {
	case startNewRoundMessage:
		err := m.user.StartNewRound(ctx)
		if err != nil {
			return fmt.Errorf("failed StartNewRound in NextRoundSuggestion: %w", err)
		}

		time.AfterFunc(time.Duration(m.user.PreferenceRoundTimeInSeconds())*time.Second, func() {
			m.concludeRoundIfUserHasNotAlready(ctx)
		})

		newMenu := NewWordGuess(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sendDefaultMessage in NextRoundSuggestion): %w", err)
		}
		return nil
	default:
		m.log.Debug("unknown command in NextRoundSuggestion", "message", message, "user_id", m.user.TelegramID())
		_, err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in NextRoundSuggestion: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
		}
		return nil
	}
}

func (m NextRoundSuggestion) concludeRoundIfUserHasNotAlready(ctx context.Context) {
	updatedUser, err := user.NewUpdatedUser(ctx, m.user)
	if err != nil {
		m.log.Error("in AfterFunc failed NewUpdatedUser: %w", err, "user_id", m.user.TelegramID())
		return
	}

	// Check if user hasn't already ended the round (with button)
	if !updatedUser.IsStillSameGuessingRound(m.user.RoundStartTime()) {
		m.log.Info("in AfterFunc round has ended before", "user_id", updatedUser.TelegramID())
		return
	}

	roundResults, err := updatedUser.ConcludeRound(ctx)
	if err != nil {
		m.log.Error("in AfterFunc failed ConcludeRound for user: %d): %w", updatedUser.TelegramID(), err)
		return
	}

	_, err = m.tgClient.SendTextMessage(
		ctx,
		updatedUser.TelegramID(),
		fmt.Sprintf("Результат раунда:\n%s", roundResults),
	)
	if err != nil {
		m.log.Error("in AfterFunc failed sending text message: %w", err)
		return
	}

	newMenu := NewRoundResult(m.tgClient, updatedUser, m.log)
	err = newMenu.sendDefaultMessage(ctx)
	if err != nil {
		m.log.Error("in AfterFunc failed sending message in WordGuess: %w", err)
		return
	}
}

func (m NextRoundSuggestion) sendDefaultMessage(ctx context.Context) error {
	teamName := m.user.CurrentTeamName()
	var msg string
	if teamName != "" {
		msg = fmt.Sprintf("Следующая команда: %s", teamName)
	} else {
		msg = "Приготовься"
	}
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		msg,
		tgTypes.KeyboardButtonsFromStrings([]string{startNewRoundMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
