package types

// MessageResponse https://core.telegram.org/bots/api#sendmessage
type MessageResponse struct {
	Ok bool `json:"ok"`
	// Exists only if Ok is false
	Description string  `json:"description,omitempty"`
	Result      Message `json:"result"`
}
