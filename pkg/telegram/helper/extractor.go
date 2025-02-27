package helper

import (
	"alias-game/pkg/telegram/types"
	"fmt"
)

func ExtractUserFromUpdate(tgUpdate types.Update) (*types.User, string, error) {
	switch {
	case tgUpdate.CallbackQuery != nil:
		return extractUserFromCallbackQuery(*tgUpdate.CallbackQuery)
	case tgUpdate.Message != nil:
		return extractUserFromMessage(*tgUpdate.Message)
	default:
		return nil, "", fmt.Errorf("unsupported Update type")
	}
}

func extractUserFromMessage(message types.Message) (*types.User, string, error) {
	if message.From == nil {
		return nil, "", fmt.Errorf("no user in Message.From")
	}
	return message.From, message.Text, nil
}

func extractUserFromCallbackQuery(callbackQuery types.CallbackQuery) (*types.User, string, error) {
	switch message := callbackQuery.Message.(type) {
	case types.Message:
		if message.Text == "" {
			return nil, "", fmt.Errorf("failed getting CallbackQuery.Message.Text")
		}
		return &callbackQuery.From, message.Text, nil
	case types.InaccessibleMessage:
		return nil, "", fmt.Errorf("InaccessibleMessage in callbackQuery")
	default:
		return nil, "", fmt.Errorf("somehow no valid message in callbackQuery")
	}
}
