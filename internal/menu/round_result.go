package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
	"slices"
	"strings"
)

const startAgainMessage = "Начать новую игру"
const currentGameResultsMessage = "Текущие результаты игры"
const endGameResultsMessage = "Результаты игры"
const nextWord2Message = "Дальше"

type RoundResult struct {
	tgClient *telegram.Client
	user     *user.User
}

func NewRoundResult(tgClient *telegram.Client, u *user.User) RoundResult {
	return RoundResult{
		tgClient: tgClient,
		user:     u,
	}
}

func (m RoundResult) Respond(ctx context.Context, message string) error {
	expectedOptions, err := m.options()
	if err != nil {
		return fmt.Errorf("error options: %w", err)
	}

	if !slices.Contains(expectedOptions, message) {
		errMessage := fmt.Sprintf("Недопустимая комманда: '%s' вместо одной из '%s'", message, strings.Join(expectedOptions, ", "))
		log.Printf("%s for user: %d in RoundResult", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in RoundResult: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in RoundResult", message)
	}

	switch message {
	case nextWord2Message:
		err := chooseNextRoundSuggestion(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseNextRoundSuggestion: %w", err)
		}
		return nil
	case currentGameResultsMessage:
		err := chooseCurrentGameResult(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseCurrentGameResult: %w", err)
		}
		return nil
	case endGameResultsMessage:
		err := chooseEndGameResult(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseEndGameResult: %w", err)
		}
		return nil
	case startAgainMessage:
		err := chooseNewStart0(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseNewStart0: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("inner logic error: unexpected answer '%s' in RoundResult respond", message)
	}
}

func chooseRoundResult(ctx context.Context, tgClient *telegram.Client, u *user.User) error {
	err := u.ChangeCurrentMenu(ctx, menuConstant.RoundResult)
	if err != nil {
		return fmt.Errorf("failed in chooseRoundResult changing current menu: %w", err)
	}
	roundResults, err := u.ConcludeRound(ctx)
	if err != nil {
		return fmt.Errorf("failed ConcludeRound for user: %d): %w", u.TelegramID(), err)
	}
	err = tgClient.SendTextMessage(
		ctx,
		u.TelegramID(),
		fmt.Sprintf("Результат раунда:\n%s", roundResults),
	)
	if err != nil {
		return fmt.Errorf("failed sending text message: %w", err)
	}
	thisMenu := NewRoundResult(tgClient, u)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseRoundResult: %w", err)
	}
	return nil
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
		return m.optionsForOneTeam(), nil
	}

	isGameEnded, err := m.user.IsGameEnded()
	if err != nil {
		return []string{}, fmt.Errorf("failed IsGameEnded for user: %d): %w", m.user.TelegramID(), err)
	}

	if isGameEnded {
		return m.optionsForMultipleTeamsGameEnded(), nil
	}
	return m.optionsForMultipleTeamsGameNotEnded(), nil
}

func (m RoundResult) optionsForOneTeam() []string {
	return []string{startAgainMessage}
}

func (m RoundResult) optionsForMultipleTeamsGameEnded() []string {
	return []string{endGameResultsMessage, startAgainMessage}
}

func (m RoundResult) optionsForMultipleTeamsGameNotEnded() []string {
	return []string{nextWord2Message, currentGameResultsMessage, startAgainMessage}
}
