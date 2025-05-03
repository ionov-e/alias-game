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

const startMessage = "Старт"

type Start0 struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

func NewStart0(tgClient *telegram.Client, u *user.User, log *slog.Logger) Start0 {
	return Start0{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (m Start0) Respond(ctx context.Context, message string) error {
	switch message {
	case startMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.SetRoundTimePredefined)
		if err != nil {
			return fmt.Errorf("failed in Start0 changing current menu: %w", err)
		}
		newMenu := NewSetRoundTimePredefined(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in Start0: %w", err)
		}
		return nil
	default:
		m.log.Debug("unknown command in Start0", "message", message, "user_id", m.user.TelegramID())
		_, err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in Start0: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s' in Start0, failed to send message: %w", message, err)
		}
		return nil
	}
}

func (m Start0) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		"Начало игры",
		tgTypes.KeyboardButtonsFromStrings([]string{startMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message in Start0: %w", err)
	}
	return nil
}
