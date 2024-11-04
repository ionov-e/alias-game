package types

// MessageOrigin https://core.telegram.org/bots/api#messageorigin
type MessageOrigin struct {
	Type string `json:"type"`
	Date uint32 `json:"date"`
}
