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

const startNewGameMessage = "Начать новую игру"

type EndGameResult struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

func NewEndGameResult(tgClient *telegram.Client, u *user.User, log *slog.Logger) EndGameResult {
	return EndGameResult{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (m EndGameResult) Respond(ctx context.Context, message string) error {
	switch message {
	case startNewGameMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.Start0)
		if err != nil {
			return fmt.Errorf("failed in EndGameResult changing current menu: %w", err)
		}
		newMenu := NewStart0(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in EndGameResult: %w", err)
		}
		return nil
	default:
		m.log.Debug("unknown command in EndGameResult", "message", message, "user_id", m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in EndGameResult: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s' in EndGameResult, failed to send message: %w", message, err)
		}
		return nil
	}
}

func (m EndGameResult) sendDefaultMessage(ctx context.Context) error {
	result := m.user.EndGameResult()
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		result,
		tgTypes.KeyboardButtonsFromStrings([]string{startNewGameMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
