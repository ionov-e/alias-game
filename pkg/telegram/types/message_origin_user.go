package types

type MessageOriginUser struct {
	Type       string `json:"type"`        // Type of the message origin, always "user"
	Date       int    `json:"date"`        // Date the message was originally sent (Unix time)
	SenderUser User   `json:"sender_user"` // User who sent the message originally
}

func (m MessageOriginUser) MessageOriginType() string { return m.Type }
func (m MessageOriginUser) MessageOriginDate() int    { return m.Date }
