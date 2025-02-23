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

type SetDictionary0 struct {
	tgClient *telegram.Client
	user     *user.User
}

const easy1DictionaryNameMessage = "Легкий словарь"
const backMessage = "Назад"

func NewSetDictionary0(tgClient *telegram.Client, u *user.User) SetDictionary0 {
	return SetDictionary0{
		tgClient: tgClient,
		user:     u,
	}
}

func (m SetDictionary0) Respond(ctx context.Context, message string) error {
	switch message {
	case easy1DictionaryNameMessage:
		err := m.user.ChooseDictionary(ctx, user.Easy1)
		if err != nil {
			return fmt.Errorf("failed ChooseDictionary in SetDictionary0: %w", err)
		}
		err = chooseSetTeamCount(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("failed chooseSetTeamCount in SetDictionary0, respond with %s: %w", message, err)
		}
		return nil
	case backMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.Start0)
		if err != nil {
			return fmt.Errorf("failed in SetDictionary0 changing current menu: %w", err)
		}
		newMenu := NewStart0(m.tgClient, m.user)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in SetDictionary0: %w", err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in SetDictionary0", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetDictionary0: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send menu message in SetDictionary0: %w", message, err)
		}
		return fmt.Errorf("unexpected message '%s' in SetDictionary0", message)
	}
}

func (m SetDictionary0) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		"Выбери набор слов",
		tgTypes.KeyboardButtonsFromStrings([]string{easy1DictionaryNameMessage, backMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message in SetDictionary: %w", err)
	}

	return nil
}
