package types

type MessageOriginChat struct {
	// Type of the message origin, always "chat"
	Type string `json:"type"`
	// Date the message was originally sent (Unix time)
	Date int `json:"date"`
	// Chat that sent the message originally
	SenderChat Chat `json:"sender_chat"`
	// Optional: Original message author signature if anonymous
	AuthorSignature string `json:"author_signature,omitempty"`
}

func (m MessageOriginChat) MessageOriginType() string { return m.Type }
func (m MessageOriginChat) MessageOriginDate() int    { return m.Date }
