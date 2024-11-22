package types

import (
	"encoding/json"
	"fmt"
)

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	// Unique identifier for this query.
	ID string `json:"id"`
	// Sender of the query.
	From User `json:"from"`
	// Optional. Message sent by the bot with the callback button that originated the query
	Message MaybeInaccessibleMessage `json:"message,omitempty"`
	// Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`
	// Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	Data string `json:"data,omitempty"`
	// Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	GameShortName string `json:"game_short_name,omitempty"`
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
