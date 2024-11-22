package types

import (
	"encoding/json"
	"fmt"
)

// SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document).
// On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
	// Identifier of the business connection for sending the message
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Optional: Identifier for the target message thread in forums
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	Voice InputFileOrString `json:"voice"`
	// Optional: Voice message caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional: Mode for parsing entities in the caption
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional: A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional: Duration of the voice message in seconds
	Duration int `json:"duration,omitempty"`
	// Optional: Send the message without sound
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional: Protect the message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional: Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional: Effect ID for adding an effect to the message
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional: Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (s SendVoice) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendVoice: %w", err)
	}
	return jsonBytes, nil
}
