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

const nextRoundMessage = "Начать раунд"

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
	case nextRoundMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.Word)
		if err != nil {
			return fmt.Errorf("failed in NextRoundSuggestion changing menu: %w", err)
		}
		newMenu := NewWordGuess(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sendDefaultMessage in NextRoundSuggestion): %w", err)
		}
		return nil
	default:
		m.log.Debug("unknown command in NextRoundSuggestion", "message", message, "user_id", m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
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
		tgTypes.KeyboardButtonsFromStrings([]string{nextRoundMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
