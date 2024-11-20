package types

import (
	"encoding/json"
	"fmt"
)

// SendMediaGroup represents a request to send a group of media items (photos, videos, documents, audios) as an album.
type SendMediaGroup struct {
	BusinessConnectionID *string          `json:"business_connection_id,omitempty"` // Optional: Identifier of the business connection for sending the message
	ChatID               interface{}      `json:"chat_id"`                          // Required: Unique identifier for the target chat or username of the target channel
	MessageThreadID      *int             `json:"message_thread_id,omitempty"`      // Optional: Identifier for the target message thread in forums
	Media                []InputMedia     `json:"media"`                            // Required: List of media items to be sent (2-10 items)
	DisableNotification  *bool            `json:"disable_notification,omitempty"`   // Optional: Send the message silently without notification sound
	ProtectContent       *bool            `json:"protect_content,omitempty"`        // Optional: Prevent the message from being forwarded or saved
	AllowPaidBroadcast   *bool            `json:"allow_paid_broadcast,omitempty"`   // Optional: Allow paid broadcasting
	MessageEffectID      *string          `json:"message_effect_id,omitempty"`      // Optional: Message effect ID for adding effects to the message
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Optional: Information about the message to reply to
	ReplyMarkup          interface{}      `json:"reply_markup,omitempty"`           // Optional: InlineKeyboardMarkup, etc.
}

func (s SendMediaGroup) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendMediaGroup: %w", err)
	}
	return jsonBytes, nil
}
