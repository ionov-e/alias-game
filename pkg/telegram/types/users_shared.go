package types

// UsersShared
// This object contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	// RequestID Identifier of the request
	RequestID int `json:"request_id"`
	// Users Information about users shared with the bot.
	Users []SharedUser `json:"users"`
}
