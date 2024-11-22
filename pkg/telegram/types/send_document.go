package types

import (
	"encoding/json"
	"fmt"
)

// SendDocument represents data needed to send a document.
type SendDocument struct {
	// Optional: Identifier of the business connection for sending the message
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Required: Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`
	// Optional: Identifier for the target message thread in forums
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// Required: Document to send as file_id, HTTP URL, or upload
	Document InputFileOrString `json:"document"`
	// Optional: Thumbnail for the document in JPEG format
	Thumbnail *InputFileOrString `json:"thumbnail,omitempty"`
	// Optional: Caption for the document, up to 1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional: Mode for parsing entities in the caption
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional: List of special entities in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional: Disable automatic content type detection
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	// Optional: Send the message without sound
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional: Protect the message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional: Allow up to 1000 messages per second with paid broadcasting
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional: Effect ID for adding an effect to the message
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional: Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional: InlineKeyboardMarkup, etc.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (s SendDocument) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendDocument: %w", err)
	}
	return jsonBytes, nil
}
