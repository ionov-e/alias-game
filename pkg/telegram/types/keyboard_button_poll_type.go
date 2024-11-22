package types

// KeyboardButtonPollType represents the type of poll allowed when the button is pressed.
type KeyboardButtonPollType struct {
	// Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
	Type string `json:"type,omitempty"`
}
