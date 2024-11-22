package types

import (
	"encoding/json"
	"fmt"
)

// SendLocation represents a request to send a point on the map.
type SendLocation struct {
	// Optional: Identifier of the business connection for sending the message
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Optional: Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// Latitude of the location
	Latitude float64 `json:"latitude"`
	// Longitude of the location
	Longitude float64 `json:"longitude"`
	// Optional: Radius of uncertainty for the location (0-1500 meters)
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// Optional: Period in seconds during which the location will be updated (see Live Locations, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod int `json:"live_period,omitempty"`
	// Optional: For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`
	// Optional: For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
	// Optional: Send the message silently without notification sound
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional: Prevent the message from being forwarded or saved
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional: Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional: Message effect ID for adding effects to the message
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional: Information about the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (s SendLocation) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendLocation: %w", err)
	}
	return jsonBytes, nil
}
