package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	"context"
	"fmt"
	"log/slog"
)

type SetTeamName struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

func NewSetTeamName(tgClient *telegram.Client, u *user.User, log *slog.Logger) SetTeamName {
	return SetTeamName{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (m SetTeamName) Respond(ctx context.Context, message string) error {
	if len(message) == 0 {
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), "Введено пустое имя")
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetTeamName: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s' in SetTeamName, failed to send message: %w", message, err)
		}
		return m.sendDefaultMessage(ctx)
	}

	if len(message) > 20 {
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), "Имя команды слишком длинное")
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetTeamName: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s' in SetTeamName, failed to send message: %w", message, err)
		}
		return m.sendDefaultMessage(ctx)
	}

	firstTeamNumberWithoutName, totalTeamCount, err := m.user.InfoForFillingTeamNames()
	if err != nil {
		return fmt.Errorf("error in user.InfoForFillingTeamNames: %w", err)
	}
	err = m.user.SetTeamName(ctx, message, firstTeamNumberWithoutName)
	if err != nil {
		return fmt.Errorf("error in user.SetTeamName: %w", err)
	}
	if firstTeamNumberWithoutName+1 == totalTeamCount {
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.SetWordCountToWinPredefined)
		if err != nil {
			return fmt.Errorf("failed in SetTeamName changing current menu: %w", err)
		}
		newMenu := NewSetWordCountToWinPredefined(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in SetTeamName: %w", err)
		}
		return nil
	}
	err = m.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("error sendDefaultMessage in SetTeamName Respond: %w", err)
	}
	return nil
}

func (m SetTeamName) sendDefaultMessage(ctx context.Context) error {
	firstTeamNumberWithoutName, _, err := m.user.InfoForFillingTeamNames()
	if err != nil {
		return fmt.Errorf("error InfoForFillingTeamNames in SetTeamName: %w", err)
	}
	err = m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Выбор названия команды (до 20 символов) для команды №%d", firstTeamNumberWithoutName+1))
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
