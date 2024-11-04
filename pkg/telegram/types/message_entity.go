package types

// MessageEntity https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type string `json:"type"`
	User User   `json:"user"`
}
