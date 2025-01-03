package helper

import (
	tgTypes "alias-game/pkg/telegram/types"
	"fmt"
)

func ExtractFromUpdate(tgUpdate tgTypes.Update) (*tgTypes.User, string, error) {
	switch {
	case tgUpdate.CallbackQuery != nil:
		return fromCallbackQuery(*tgUpdate.CallbackQuery)
	case tgUpdate.Message != nil:
		return fromMessage(*tgUpdate.Message)
	default:
		return nil, "", fmt.Errorf("unsupported Update type: %+v", tgUpdate)
	}
}

func fromMessage(message tgTypes.Message) (*tgTypes.User, string, error) {
	if message.From == nil {
		return nil, "", fmt.Errorf("no user in Message.From: %+v", message)
	}
	return message.From, message.Text, nil
}

func fromCallbackQuery(callbackQuery tgTypes.CallbackQuery) (*tgTypes.User, string, error) {
	switch message := callbackQuery.Message.(type) {
	case tgTypes.Message:
		if message.Text == "" {
			return &tgTypes.User{}, "", fmt.Errorf("failed getting CallbackQuery.Message.Text: %+v", callbackQuery)
		}
		return &callbackQuery.From, message.Text, nil
	case tgTypes.InaccessibleMessage:
		return &tgTypes.User{}, "", fmt.Errorf("InaccessibleMessage in callbackQuery: %+v", callbackQuery)
	default:
		return &tgTypes.User{}, "", fmt.Errorf("somehow no valid message in callbackQuery: %+v", callbackQuery)
	}
}
