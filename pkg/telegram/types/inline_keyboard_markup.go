package types

// InlineKeyboardMarkup represents an inline keyboard that appears with the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an array of InlineKeyboardButton objects.
}
