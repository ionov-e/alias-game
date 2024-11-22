package types

// ReplyKeyboardMarkup
// This object represents a custom keyboard with reply options. Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]KeyboardButton `json:"keyboard"`
	// Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false.
	IsPersistent bool `json:"is_persistent,omitempty"`
	// Optional. Requests clients to resize the keyboard vertically for optimal fit. Defaults to false.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	// Optional. Requests clients to hide the keyboard as soon as it's been used. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	// Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters.
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	// Optional. Use this parameter to show the keyboard to specific users only.
	Selective bool `json:"selective,omitempty"`
}

// PaidMediaType implements the ReplyKeyboardMarkup interface
func (r ReplyKeyboardMarkup) PaidMediaType() string {
	return "ReplyKeyboardMarkup"
}
