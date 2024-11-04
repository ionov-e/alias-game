package types

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID uint64 `json:"update_id"` //nolint:tagliatelle
	// Optional: new incoming message of any kind - text, photo, sticker, etc.
	Message Message `json:"message"`
	// Optional: new message from a connected business account
	BusinessMessage Message `json:"business_message"` //nolint:tagliatelle
}
