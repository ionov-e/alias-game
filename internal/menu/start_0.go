package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
)

const startMessage = "Старт"

type Start0 struct {
	tgClient *telegram.Client
	user     *user.User
}

func NewStart0(tgClient *telegram.Client, user *user.User) Start0 {
	return Start0{
		tgClient: tgClient,
		user:     user,
	}
}

func (m Start0) Respond(ctx context.Context, message string) error {
	switch message {
	case startMessage:
		err := chooseSetRoundTime(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseDictionaryChoice0: %w", err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная команда: '%s'", message)
		log.Printf("%s for user: %d in Start0", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in Start0: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in Start0", message)
	}
}

func chooseNewStart0(ctx context.Context, client *telegram.Client, user *user.User) error {
	err := user.ChangeCurrentMenu(ctx, menuConstant.Start0)
	if err != nil {
		return fmt.Errorf("failed in chooseNewStart0 changing current menu: %w", err)
	}
	thisMenu := NewStart0(client, user)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseNewStart0: %w", err)
	}
	return nil
}

func (m Start0) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		"Начало игры",
		tgTypes.KeyboardButtonsFromStrings(m.options()),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}

func (m Start0) options() []string {
	return []string{
		startMessage,
	}
}
