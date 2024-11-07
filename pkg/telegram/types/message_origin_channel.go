package types

type MessageOriginChannel struct {
	Type            string  `json:"type"`                       // Type of the message origin, always "channel"
	Date            int     `json:"date"`                       // Date the message was originally sent (Unix time)
	Chat            Chat    `json:"chat"`                       // Channel chat where message was sent
	MessageID       int     `json:"message_id"`                 // Unique message ID in the chat
	AuthorSignature *string `json:"author_signature,omitempty"` // Optional: Signature of original post author
}

func (m MessageOriginChannel) MessageOriginType() string { return m.Type }
func (m MessageOriginChannel) MessageOriginDate() int    { return m.Date }
