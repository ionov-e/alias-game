package menu

import (
	dbConstants "alias-game/internal/database/constants"
	aliasUser "alias-game/internal/service/alias/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
)

type DictionaryChoice0 struct {
	client *telegram.Client
	user   *aliasUser.User
}

const DictionaryChoice0Name = "Выбор слов"
const Easy1Name = "Легкий словарь"

func NewDictionaryChoice0(client *telegram.Client, user *aliasUser.User) Menu {
	return &DictionaryChoice0{
		client: client,
		user:   user,
	}
}

func ChooseDictionaryChoice0(ctx context.Context, client *telegram.Client, user *aliasUser.User) error {
	thisMenu := NewDictionaryChoice0(client, user)
	err := user.ChangeCurrentMenu(ctx, thisMenu.RedisKey())
	if err != nil {
		return fmt.Errorf("failed in ChooseDictionaryChoice0 changing current menu: %w", err)
	}
	err = thisMenu.Message(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in ChooseNewStart0 in ChooseDictionaryChoice0: %w", err)
	}

	return nil
}

func (m DictionaryChoice0) Message(ctx context.Context) error {
	err := m.client.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		"Выбери набор слов",
		tgTypes.KeyboardButtonsFromStrings(m.Options()),
	)
	if err != nil {
		return fmt.Errorf("failed sending message in DictionaryChoice0: %w", err)
	}

	return nil
}

func (m DictionaryChoice0) RedisKey() dbConstants.MenuKeyStored {
	return dbConstants.MenuDictionaryChoice0Key
}

func (m DictionaryChoice0) Options() []string {
	return []string{
		Easy1Name,
		BackString,
	}
}

func (m DictionaryChoice0) Respond(ctx context.Context, message string) error {
	switch message {
	case Easy1Name:
		err := m.user.ChooseDictionary(ctx, dbConstants.Easy1)
		if err != nil {
			return fmt.Errorf("failed ChooseDictionary in DictionaryChoice0: %w", err)
		}
		err = ChooseWord(ctx, 0, m.client, m.user)
		if err != nil {
			return fmt.Errorf("failed DictionaryChoice0 Respond with %s: %w", message, err)
		}
		return nil
	case BackString:
		err := ChooseNewStart0(ctx, m.client, m.user)
		if err != nil {
			return fmt.Errorf("failed DictionaryChoice0 Respond with %s: %w", message, err)
		}
		return nil
	default:
		err := m.Message(ctx)
		if err != nil {
			return fmt.Errorf("не ожиданный ответ '%s', провалились отправить сообщение: %w", message, err)
		}
		return fmt.Errorf("неожиданный ответ '%s' из меню выбора словаря DictionaryChoice0", message)
	}
}
