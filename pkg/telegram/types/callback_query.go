package types

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	ID              string                    `json:"id"`                          // Unique identifier for this query.
	From            User                      `json:"from"`                        // Sender of the query.
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query.
	InlineMessageID *string                   `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode.
	ChatInstance    string                    `json:"chat_instance"`               // Global identifier for the chat associated with the callback button.
	Data            *string                   `json:"data,omitempty"`              // Optional. Data associated with the callback button.
	GameShortName   *string                   `json:"game_short_name,omitempty"`   // Optional. Short name of a game to be returned.
}
