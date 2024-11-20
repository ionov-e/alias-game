package types

import (
	"encoding/json"
	"fmt"
)

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	ID              string                   `json:"id"`                          // Unique identifier for this query.
	From            User                     `json:"from"`                        // Sender of the query.
	Message         MaybeInaccessibleMessage `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query.
	InlineMessageID *string                  `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode.
	ChatInstance    string                   `json:"chat_instance"`               // Global identifier for the chat associated with the callback button.
	Data            *string                  `json:"data,omitempty"`              // Optional. Data associated with the callback button.
	GameShortName   *string                  `json:"game_short_name,omitempty"`   // Optional. Short name of a game to be returned.
}

func (cq *CallbackQuery) UnmarshalJSON(data []byte) error {
	type Alias CallbackQuery
	aux := struct {
		Message json.RawMessage `json:"message"`
		*Alias
	}{
		Alias: (*Alias)(cq),
	}

	// Unmarshal into the auxiliary struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("failed to unmarshal CallbackQuery: %w", err)
	}

	// Process the "message" field specifically
	if len(aux.Message) > 0 {
		var msg Message
		if err := json.Unmarshal(aux.Message, &msg); err == nil {
			cq.Message = msg
			return nil
		}

		var im InaccessibleMessage
		if err := json.Unmarshal(aux.Message, &im); err == nil {
			cq.Message = im
			return nil
		}

		return fmt.Errorf("failed to unmarshal Message in CallbackQuery")
	}

	return nil
}
