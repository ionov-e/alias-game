package types

import (
	"encoding/json"
)

// SendVideoNote represents a request to send a video note.
type SendVideoNote struct {
	BusinessConnectionID *string          `json:"business_connection_id,omitempty"` // Optional: Identifier of the business connection for sending the message
	ChatID               interface{}      `json:"chat_id"`                          // Required: Unique identifier for the target chat or username of the target channel
	MessageThreadID      *int             `json:"message_thread_id,omitempty"`      // Optional: Identifier for the target message thread in forums
	VideoNote            interface{}      `json:"video_note"`                       // Required: Video note to send as file_id, HTTP URL, or upload
	Duration             *int             `json:"duration,omitempty"`               // Optional: Duration of the video in seconds
	Length               *int             `json:"length,omitempty"`                 // Optional: Diameter of the video message (width & height)
	Thumbnail            *interface{}     `json:"thumbnail,omitempty"`              // Optional: Thumbnail for the video note (JPEG, <200kB, 320px max width/height)
	DisableNotification  *bool            `json:"disable_notification,omitempty"`   // Optional: Send the message silently without notification sound
	ProtectContent       *bool            `json:"protect_content,omitempty"`        // Optional: Prevent the message from being forwarded or saved
	AllowPaidBroadcast   *bool            `json:"allow_paid_broadcast,omitempty"`   // Optional: Allow paid broadcasting
	MessageEffectID      *string          `json:"message_effect_id,omitempty"`      // Optional: Message effect ID for adding effects to the message
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Optional: Information about the message to reply to
	ReplyMarkup          interface{}      `json:"reply_markup,omitempty"`           // Optional: InlineKeyboardMarkup, etc.

}

func (s *SendVideoNote) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
