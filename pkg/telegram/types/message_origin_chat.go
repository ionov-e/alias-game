package types

type MessageOriginChat struct {
	Type            string  `json:"type"`                       // Type of the message origin, always "chat"
	Date            int     `json:"date"`                       // Date the message was originally sent (Unix time)
	SenderChat      Chat    `json:"sender_chat"`                // Chat that sent the message originally
	AuthorSignature *string `json:"author_signature,omitempty"` // Optional: Original message author signature if anonymous
}

func (m MessageOriginChat) MessageOriginType() string { return m.Type }
func (m MessageOriginChat) MessageOriginDate() int    { return m.Date }
