package types

// LoginUrl represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginUrl struct {
	// An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in https://core.telegram.org/widgets/login#receiving-authorization-data
	URL string `json:"url"`
	// Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`
	// Optional. Username of a bot, which will be used for user authorization. See https://core.telegram.org/widgets/login#setting-up-a-bot. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See https://core.telegram.org/widgets/login#receiving-authorization-data
	BotUsername string `json:"bot_username,omitempty"`
	// Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}
