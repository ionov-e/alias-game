package types

type MessageOriginUser struct {
	// Type of the message origin, always "user"
	Type string `json:"type"`
	// Date the message was originally sent (Unix time)
	Date int `json:"date"`
	// User who sent the message originally
	SenderUser User `json:"sender_user"`
}

func (m MessageOriginUser) MessageOriginType() string { return m.Type }
func (m MessageOriginUser) MessageOriginDate() int    { return m.Date }
