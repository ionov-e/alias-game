package types

import (
	"encoding/json"
)

// SendLocation represents a request to send a point on the map.
type SendLocation struct {
	BusinessConnectionID *string          `json:"business_connection_id,omitempty"` // Optional: Identifier of the business connection for sending the message
	ChatID               interface{}      `json:"chat_id"`                          // Required: Unique identifier for the target chat or username of the target channel
	MessageThreadID      *int             `json:"message_thread_id,omitempty"`      // Optional: Identifier for the target message thread in forums
	Latitude             float64          `json:"latitude"`                         // Required: Latitude of the location
	Longitude            float64          `json:"longitude"`                        // Required: Longitude of the location
	HorizontalAccuracy   *float64         `json:"horizontal_accuracy,omitempty"`    // Optional: Radius of uncertainty for the location (0-1500 meters)
	LivePeriod           *int             `json:"live_period,omitempty"`            // Optional: Period in seconds to update the live location
	Heading              *int             `json:"heading,omitempty"`                // Optional: Direction of movement for live location
	ProximityAlertRadius *int             `json:"proximity_alert_radius,omitempty"` // Optional: Maximum distance for proximity alerts
	DisableNotification  *bool            `json:"disable_notification,omitempty"`   // Optional: Send the message silently without notification sound
	ProtectContent       *bool            `json:"protect_content,omitempty"`        // Optional: Prevent the message from being forwarded or saved
	AllowPaidBroadcast   *bool            `json:"allow_paid_broadcast,omitempty"`   // Optional: Allow paid broadcasting
	MessageEffectID      *string          `json:"message_effect_id,omitempty"`      // Optional: Message effect ID for adding effects to the message
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Optional: Information about the message to reply to
	ReplyMarkup          interface{}      `json:"reply_markup,omitempty"`           // Optional: InlineKeyboardMarkup, etc.
}

// ToJSON converts the SendLocation struct to a JSON string.
func (s *SendLocation) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
