package telegram_type

import (
	tgTypes "alias-game/pkg/telegram/types"
	"fmt"
)

func ExtractUserFromUpdate(tgUpdate tgTypes.Update) (*tgTypes.User, string, error) {
	switch {
	case tgUpdate.CallbackQuery != nil:
		return extractUserFromCallbackQuery(*tgUpdate.CallbackQuery)
	case tgUpdate.Message != nil:
		return extractUserFromMessage(*tgUpdate.Message)
	default:
		return nil, "", fmt.Errorf("unsupported Update type")
	}
}

func extractUserFromMessage(message tgTypes.Message) (*tgTypes.User, string, error) {
	if message.From == nil {
		return nil, "", fmt.Errorf("no user in Message.From")
	}
	return message.From, message.Text, nil
}

func extractUserFromCallbackQuery(callbackQuery tgTypes.CallbackQuery) (*tgTypes.User, string, error) {
	switch message := callbackQuery.Message.(type) {
	case tgTypes.Message:
		if message.Text == "" {
			return nil, "", fmt.Errorf("failed getting CallbackQuery.Message.Text")
		}
		return &callbackQuery.From, message.Text, nil
	case tgTypes.InaccessibleMessage:
		return nil, "", fmt.Errorf("InaccessibleMessage in callbackQuery")
	default:
		return nil, "", fmt.Errorf("somehow no valid message in callbackQuery")
	}
}
