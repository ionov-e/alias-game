package menu

import (
	dbConstants "alias-game/internal/database/constants"
	aliasUser "alias-game/internal/service/alias/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
)

type Start0 struct {
	client *telegram.Client
	user   *aliasUser.User
}

func NewStart0(client *telegram.Client, user *aliasUser.User) Menu {
	return &Start0{
		client: client,
		user:   user,
	}
}

func ChooseNewStart0(ctx context.Context, client *telegram.Client, user *aliasUser.User) error {
	err := user.ChangeCurrentMenu(ctx, dbConstants.MenuStart0Key)
	if err != nil {
		return fmt.Errorf("failed in ChooseNewStart0 changing current menu: %w", err)
	}
	thisMenu := NewStart0(client, user)
	err = thisMenu.DefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in ChooseNewStart0: %w", err)
	}
	return nil
}

func (m Start0) DefaultMessage(ctx context.Context) error {
	err := m.client.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		"Начало игры",
		tgTypes.KeyboardButtonsFromStrings(m.Options()),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}

func (m Start0) Options() []string {
	return []string{
		DictionaryChoice0Name,
	}
}

func (m Start0) Respond(ctx context.Context, message string) error {
	switch message {
	case DictionaryChoice0Name:
		err := ChooseDictionaryChoice0(ctx, m.client, m.user)
		if err != nil {
			return fmt.Errorf("error ChooseDictionaryChoice0: %w", err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in Start0", errMessage, m.user.TelegramID())
		err := m.client.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in Start0: %w", message, err)
		}
		err = m.DefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in Start0", message)
	}
}
