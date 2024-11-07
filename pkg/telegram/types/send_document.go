package types

import (
	"encoding/json"
)

// SendDocument represents data needed to send a document.
type SendDocument struct {
	BusinessConnectionID        *string          `json:"business_connection_id,omitempty"`         // Optional: Identifier of the business connection for sending the message
	ChatID                      interface{}      `json:"chat_id"`                                  // Required: Unique identifier for the target chat or username of the target channel
	MessageThreadID             *int             `json:"message_thread_id,omitempty"`              // Optional: Identifier for the target message thread in forums
	Document                    interface{}      `json:"document"`                                 // Required: Document to send as file_id, HTTP URL, or upload
	Thumbnail                   interface{}      `json:"thumbnail,omitempty"`                      // Optional: Thumbnail for the document in JPEG format
	Caption                     *string          `json:"caption,omitempty"`                        // Optional: Caption for the document, up to 1024 characters
	ParseMode                   *string          `json:"parse_mode,omitempty"`                     // Optional: Mode for parsing entities in the caption
	CaptionEntities             *[]MessageEntity `json:"caption_entities,omitempty"`               // Optional: List of special entities in the caption
	DisableContentTypeDetection *bool            `json:"disable_content_type_detection,omitempty"` // Optional: Disable automatic content type detection
	DisableNotification         *bool            `json:"disable_notification,omitempty"`           // Optional: Send the message without sound
	ProtectContent              *bool            `json:"protect_content,omitempty"`                // Optional: Protect the message from forwarding and saving
	AllowPaidBroadcast          *bool            `json:"allow_paid_broadcast,omitempty"`           // Optional: Allow up to 1000 messages per second with paid broadcasting
	MessageEffectID             *string          `json:"message_effect_id,omitempty"`              // Optional: Effect ID for adding an effect to the message
	ReplyParameters             *ReplyParameters `json:"reply_parameters,omitempty"`               // Optional: Description of the message to reply to
	ReplyMarkup                 interface{}      `json:"reply_markup,omitempty"`                   // Optional: InlineKeyboardMarkup, etc.
}

func (s *SendDocument) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
