package types

type MessageOriginChannel struct {
	// Type of the message origin, always "channel"
	Type string `json:"type"`
	// Date the message was originally sent (Unix time)
	Date int `json:"date"`
	// Channel chat where message was sent
	Chat Chat `json:"chat"`
	// Unique message ID in the chat
	MessageID int `json:"message_id"`
	// Optional: Signature of original post author
	AuthorSignature string `json:"author_signature,omitempty"`
}

func (m MessageOriginChannel) MessageOriginType() string { return m.Type }
func (m MessageOriginChannel) MessageOriginDate() int    { return m.Date }
