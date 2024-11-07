package types

// LoginUrl represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginUrl struct {
	URL                string  `json:"url"`                            // HTTPS URL to be opened with user authorization data added to the query string.
	ForwardText        *string `json:"forward_text,omitempty"`         // Optional. New text of the button in forwarded messages.
	BotUsername        *string `json:"bot_username,omitempty"`         // Optional. Username of a bot for user authorization.
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"` // Optional. True to request permission for the bot to send messages to the user.
}
