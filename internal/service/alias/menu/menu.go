package menu

import (
	dbConstants "alias-game/internal/database/constants"
	aliasUser "alias-game/internal/service/alias/user"
	"alias-game/pkg/telegram"
	"context"
	"fmt"
)

const BackString = "Назад"

// Menu defines where user is currently at. In bot menu/settings, or maybe currently answering a question
type Menu interface {
	RedisKey() dbConstants.MenuKeyStored
	Options() []string
	Respond(ctx context.Context, message string) error
	DefaultMessage(ctx context.Context) error
}

func FactoryMethod(
	menuKeyString string,
	client *telegram.Client,
	user *aliasUser.User,
) (Menu, error) {
	menuKey := dbConstants.MenuKeyStored(menuKeyString)

	if menuKey.IsWord() {
		wordNumber, err := menuKey.WordNumber()
		if err != nil {
			return nil, fmt.Errorf("failed getting word number from '%s': %w", menuKeyString, err)
		}
		return NewWord(wordNumber, client, user), nil
	}

	switch menuKey {
	case dbConstants.Empty:
	case dbConstants.MenuStart0Key:
		return NewStart0(client, user), nil
	case dbConstants.MenuDictionaryChoice0Key:
		return NewDictionaryChoice0(client, user), nil
	}
	return nil, fmt.Errorf("menu factory called with: '%s' - no match", menuKeyString)
}
