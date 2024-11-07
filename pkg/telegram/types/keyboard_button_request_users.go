package types

// KeyboardButtonRequestUsers defines criteria to request suitable users.

type KeyboardButtonRequestUsers struct {
	RequestID       int32 `json:"request_id"`                 // Signed 32-bit identifier of the request, unique within the message
	UserIsBot       *bool `json:"user_is_bot,omitempty"`      // Optional. True to request bots, False for regular users
	UserIsPremium   *bool `json:"user_is_premium,omitempty"`  // Optional. True to request premium users, False for non-premium users
	MaxQuantity     *int  `json:"max_quantity,omitempty"`     // Optional. Maximum number of users to be selected; 1-10. Defaults to 1
	RequestName     *bool `json:"request_name,omitempty"`     // Optional. True to request the users' first and last names
	RequestUsername *bool `json:"request_username,omitempty"` // Optional. True to request the users' usernames
	RequestPhoto    *bool `json:"request_photo,omitempty"`    // Optional. True to request the users' photos
}
