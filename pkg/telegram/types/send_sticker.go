package types

import (
	"encoding/json"
	"fmt"
)

// SendSticker - method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
type SendSticker struct {
	// Optional. Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files ».
	Sticker InputFileOrString `json:"sticker"`
	// Optional. Emoji associated with the sticker; only for just uploaded stickers
	Emoji string `json:"emoji,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional. Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional. InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (s SendSticker) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendSticker: %w", err)
	}
	return jsonBytes, nil
}
