package helper

import (
	menuConstant "alias-game/internal/constant/menu"
	userEntity "alias-game/internal/entity/user"
	"alias-game/internal/service/responder/menu"
	"alias-game/pkg/telegram"
	"context"
	"fmt"
)

type MenuServiceInterface interface {
	Respond(ctx context.Context, message string) error
}

func MenuFactory(tgClient *telegram.Client, user *userEntity.User) (MenuServiceInterface, error) {
	menuKeyString := user.CurrentMenuKey()
	menuKey := menuConstant.Key(menuKeyString)

	if menuKey.IsWord() {
		wordNumber, err := menuKey.WordNumber()
		if err != nil {
			return nil, fmt.Errorf("failed getting word number from '%s': %w", menuKeyString, err)
		}
		return menu.NewWordGuess(wordNumber, &tgClient, user), nil
	}

	switch menuKey {
	case menuConstant.Start0Key:
		return menu.NewStart0(&tgClient, user), nil
	case menuConstant.DictionaryChoice0Key:
		return menu.NewDictionaryChoice0(&tgClient, user), nil
	}
	return nil, fmt.Errorf("menu factory called with: '%s' - no match", menuKeyString)
}
