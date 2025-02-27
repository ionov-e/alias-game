package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log/slog"
	"slices"
)

const startAgainMessage = "Начать новую игру"
const currentGameResultsMessage = "Текущие результаты игры"
const endGameResultsMessage = "Результаты игры"
const nextWord2Message = "Дальше"

type RoundResult struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

func NewRoundResult(tgClient *telegram.Client, u *user.User, log *slog.Logger) RoundResult {
	return RoundResult{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (m RoundResult) Respond(ctx context.Context, message string) error {
	expectedOptions, err := m.options()
	if err != nil {
		return fmt.Errorf("error options: %w", err)
	}

	if !slices.Contains(expectedOptions, message) {
		m.log.Debug(
			"unknown command in RoundResult",
			slog.String("message", message),
			slog.Any("expected", expectedOptions),
			slog.Int64("user_id", m.user.TelegramID()),
		)
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in RoundResult: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
		}
		return nil
	}

	switch message {
	case nextWord2Message:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.NextRoundSuggestion)
		if err != nil {
			return fmt.Errorf("failed in RoundResult changing current menu: %w", err)
		}
		newMenu := NewNextRoundSuggestion(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in RoundResult: %w", err)
		}
		return nil
	case currentGameResultsMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.CurrentGameResult)
		if err != nil {
			return fmt.Errorf("failed in RoundResult changing current menu: %w", err)
		}
		newMenu := NewCurrentGameResult(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in RoundResult: %w", err)
		}
		return nil
	case endGameResultsMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.EndGameResult)
		if err != nil {
			return fmt.Errorf("failed in RoundResult changing current menu: %w", err)
		}
		newMenu := NewEndGameResult(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in RoundResult: %w", err)
		}
		err = m.user.ClearGame(ctx)
		if err != nil {
			return fmt.Errorf("failed clearing game: %w", err)
		}
		return nil
	case startAgainMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.Start0)
		if err != nil {
			return fmt.Errorf("failed in RoundResult changing current menu: %w", err)
		}
		newMenu := NewStart0(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in RoundResult: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("inner logic error: unexpected answer '%s' in RoundResult respond", message)
	}
}

func (m RoundResult) sendDefaultMessage(ctx context.Context) error {
	options, err := m.options()
	if err != nil {
		return fmt.Errorf("failed get options for user: %d): %w", m.user.TelegramID(), err)
	}

	err = m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		"Выбери действие",
		tgTypes.KeyboardButtonsFromStrings(options),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}

func (m RoundResult) options() ([]string, error) {
	if m.user.AllTeamsCount() == 1 {
		return []string{startAgainMessage}, nil
	} // One Team

	isGameEnded, err := m.user.IsGameEnded()
	if err != nil {
		return []string{}, fmt.Errorf("failed IsGameEnded for user: %d): %w", m.user.TelegramID(), err)
	}

	if isGameEnded {
		return []string{endGameResultsMessage}, nil // Multiple Teams: Game Ended
	}
	return []string{nextWord2Message, currentGameResultsMessage, startAgainMessage}, nil // Multiple Teams: Game Not Ended
}
