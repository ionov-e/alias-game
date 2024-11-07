package types

// SharedUser
// This object contains information about a user that was shared with the bot using a KeyboardButtonRequestUsers button.
type SharedUser struct {
	// UserID Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers.
	UserID int64 `json:"user_id"`
	// FirstName Optional. First name of the user, if the name was requested by the bot
	FirstName *string `json:"first_name,omitempty"`
	// LastName Optional. Last name of the user, if the name was requested by the bot
	LastName *string `json:"last_name,omitempty"`
	// Username Optional. Username of the user, if the username was requested by the bot
	Username *string `json:"username,omitempty"`
	// Photo Optional. Available sizes of the chat photo, if the photo was requested by the bot
	Photo []PhotoSize `json:"photo,omitempty"`
}
