package types

// KeyboardButtonRequestUsers defines the criteria used to request suitable users. Information about the selected users will be shared with the bot when the corresponding button is pressed.
type KeyboardButtonRequestUsers struct {
	// Signed 32-bit identifier of the request, unique within the message
	RequestID int32 `json:"request_id"`
	// Optional. True to request bots, False for regular users
	UserIsBot bool `json:"user_is_bot,omitempty"`
	// Optional. True to request premium users, False for non-premium users
	UserIsPremium bool `json:"user_is_premium,omitempty"`
	// Optional. Maximum number of users to be selected; 1-10. Defaults to 1
	MaxQuantity int `json:"max_quantity,omitempty"`
	// Optional. True to request the users' first and last names
	RequestName bool `json:"request_name,omitempty"`
	// Optional. True to request the users' usernames
	RequestUsername bool `json:"request_username,omitempty"`
	// Optional. True to request the users' photos
	RequestPhoto bool `json:"request_photo,omitempty"`
}
