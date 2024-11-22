package types

import (
	"encoding/json"
	"fmt"
)

// SendMediaGroup represents a request to send a group of media items (photos, videos, documents, audios) as an album.
type SendMediaGroup struct {
	// Optional: Identifier of the business connection for sending the message
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Optional: Identifier for the target message thread in forums
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// Array of InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo. A JSON-serialized array describing messages to be sent, must include 2-10 items
	Media []InputMedia `json:"media"`
	// Optional: Send the message silently without notification sound
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional: Prevent the message from being forwarded or saved
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional: Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional: Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional: Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
}

func (s SendMediaGroup) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendMediaGroup: %w", err)
	}
	return jsonBytes, nil
}
