package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	userConstant "alias-game/internal/constant/user"
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
		err := m.user.ChooseDictionary(ctx, userConstant.Easy1)
		if err != nil {
			return fmt.Errorf("failed ChooseDictionary in SetDictionary: %w", err)
		}
		err = chooseSetTeamCount(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("failed chooseSetTeamCount respond with %s: %w", message, err)
		}
		return nil
	case backMessage:
		err := chooseNewStart0(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("failed SetDictionary respond with %s: %w", message, err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in SetDictionary", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetDictionary: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send menu message in SetDictionary: %w", message, err)
		}
		return fmt.Errorf("unexpected message '%s' in SetDictionary", message)
	}
}

// TODO get rid of these methods
func chooseDictionaryChoice0(ctx context.Context, client *telegram.Client, u *user.User) error {
	err := u.ChangeCurrentMenu(ctx, menuConstant.SetDictionary)
	if err != nil {
		return fmt.Errorf("failed in chooseDictionaryChoice0 changing current menu: %w", err)
	}
	thisMenu := NewSetDictionary0(client, u)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseDictionaryChoice0: %w", err)
	}
	return nil
}

func (m SetDictionary0) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		"Выбери набор слов",
		tgTypes.KeyboardButtonsFromStrings(m.options()),
	)
	if err != nil {
		return fmt.Errorf("failed sending message in SetDictionary: %w", err)
	}

	return nil
}

func (m SetDictionary0) options() []string {
	return []string{
		easy1DictionaryNameMessage,
		backMessage,
	}
}
