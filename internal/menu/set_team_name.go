package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	"context"
	"fmt"
)

type SetTeamName struct {
	tgClient *telegram.Client
	user     *user.User
}

func NewSetTeamName(tgClient *telegram.Client, user *user.User) SetTeamName {
	return SetTeamName{
		tgClient: tgClient,
		user:     user,
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
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
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
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
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
		err = chooseSetWordCountToWinPredefined(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseSetWordCountToWinPredefined: %w", err)
		}
		return nil
	}
	err = m.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("error sendDefaultMessage in SetTeamName Respond: %w", err)
	}
	return nil
}

func chooseSetTeamName(ctx context.Context, client *telegram.Client, user *user.User) error {
	err := user.ChangeCurrentMenu(ctx, menuConstant.SetTeamName)
	if err != nil {
		return fmt.Errorf("failed in chooseSetTeamNameChoice changing current menu: %w", err)
	}
	thisMenu := NewSetTeamName(client, user)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseSetTeamNameChoice: %w", err)
	}
	return nil
}

func (m SetTeamName) sendDefaultMessage(ctx context.Context) error {
	firstTeamNumberWithoutName, _, err := m.user.InfoForFillingTeamNames()
	if err != nil {
		return fmt.Errorf("error InfoForFillingTeamNames: %w", err)
	}
	err = m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Выбор названия команды (до 20 символов) для команды №%d", firstTeamNumberWithoutName+1))
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
