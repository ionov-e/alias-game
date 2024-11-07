package types

// KeyboardButton represents a button of the reply keyboard.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed.
	Text string `json:"text"`
	// Optional. If specified, pressing the button will open a list of suitable users. Available in private chats only.
	RequestUsers *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	// Optional. If specified, pressing the button will open a list of suitable chats. Available in private chats only.
	RequestChat *KeyboardButtonRequestChat `json:"request_chat,omitempty"`
	// Optional. If true, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	RequestContact *bool `json:"request_contact,omitempty"`
	// Optional. If true, the user's current location will be sent when the button is pressed. Available in private chats only.
	RequestLocation *bool `json:"request_location,omitempty"`
	// Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
	// Optional. If specified, the described Web App will be launched when the button is pressed.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}
