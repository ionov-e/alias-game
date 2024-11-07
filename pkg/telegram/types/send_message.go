package types

import "encoding/json"

// SendMessage
// Use this method to send text messages. On success, the sent Message is returned.
type SendMessage struct {
	// Optional. Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID *string `json:"business_connection_id,omitempty"`
	// Required. Unique identifier for the target chat or username of the target channel
	ChatID int64 `json:"chat_id"` // Either Integer or String
	// Optional. Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadID *int `json:"message_thread_id,omitempty"`
	// Required. Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text"`
	// Optional. Mode for parsing entities in the message text.
	ParseMode *string `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in message text
	Entities []MessageEntity `json:"entities,omitempty"`
	// Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Optional. Pass True to allow up to 1000 messages per second for a fee of 0.1 Telegram Stars per message
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Optional. Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectID *string `json:"message_effect_id,omitempty"`
	// Optional. Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional. Additional interface options
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // Can be InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply
}

func (s SendMessage) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
