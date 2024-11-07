package types

// ReplyKeyboardRemove requests clients to remove the custom keyboard and display the default letter-keyboard.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`     // Requests clients to remove the custom keyboard.
	Selective      *bool `json:"selective,omitempty"` // Optional. If true, removes the keyboard for specific users only.
}
