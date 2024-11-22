package types

import (
	"encoding/json"
	"fmt"
)

// SendPhoto represents a request to send a photo message.
type SendPhoto struct {
	// Optional: Identifier of the business connection for sending the message
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`
	// Optional: Identifier for the target message thread in forums
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More: https://core.telegram.org/bots/api#sending-files
	Photo InputFileOrString `json:"photo"`
	// Optional: Caption for the photo, up to 1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional: Mode for parsing entities in the caption
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional: A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional: Display caption above the media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`
	// Optional: Display the photo with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
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

func (s SendPhoto) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendPhoto: %w", err)
	}
	return jsonBytes, nil
}
