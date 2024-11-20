package types

import (
	"encoding/json"
	"fmt"
)

// SendVoice represents a request to send a voice message.
type SendVoice struct {
	BusinessConnectionID *string          `json:"business_connection_id,omitempty"` // Optional: Identifier of the business connection for sending the message
	ChatID               interface{}      `json:"chat_id"`                          // Required: Unique identifier for the target chat or username of the target channel
	MessageThreadID      *int             `json:"message_thread_id,omitempty"`      // Optional: Identifier for the target message thread in forums
	Voice                interface{}      `json:"voice"`                            // Required: Voice file to send as file_id, HTTP URL, or upload
	Caption              *string          `json:"caption,omitempty"`                // Optional: Caption for the voice message, up to 1024 characters
	ParseMode            *string          `json:"parse_mode,omitempty"`             // Optional: Mode for parsing entities in the caption
	CaptionEntities      *[]MessageEntity `json:"caption_entities,omitempty"`       // Optional: List of special entities in the caption
	Duration             *int             `json:"duration,omitempty"`               // Optional: Duration of the voice message in seconds
	DisableNotification  *bool            `json:"disable_notification,omitempty"`   // Optional: Send the message without sound
	ProtectContent       *bool            `json:"protect_content,omitempty"`        // Optional: Protect the message from forwarding and saving
	AllowPaidBroadcast   *bool            `json:"allow_paid_broadcast,omitempty"`   // Optional: Allow up to 1000 messages per second with paid broadcasting
	MessageEffectID      *string          `json:"message_effect_id,omitempty"`      // Optional: Effect ID for adding an effect to the message
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Optional: Description of the message to reply to
	ReplyMarkup          interface{}      `json:"reply_markup,omitempty"`           // Optional: InlineKeyboardMarkup, etc.
}

func (s SendVoice) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendVoice: %w", err)
	}
	return jsonBytes, nil
}
