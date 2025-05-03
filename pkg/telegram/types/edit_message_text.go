package types

import (
	"encoding/json"
	"fmt"
)

// EditMessageText
// Use this method to edit text and game messages.
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
// https://core.telegram.org/bots/api#editmessagetext
type EditMessageText struct {
	// Optional. Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID int64 `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`
	// Optional. Mode for parsing entities in the message text. See formatting options for more details: https://core.telegram.org/bots/api#formatting-options
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty"`
	// Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	// Optional. InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // Can be InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply
}

func (s EditMessageText) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling EditMessageText: %w", err)
	}
	return jsonBytes, nil
}
