package types

type MessageOriginHiddenUser struct {
	Type           string `json:"type"`             // Type of the message origin, always "hidden_user"
	Date           int    `json:"date"`             // Date the message was originally sent (Unix time)
	SenderUserName string `json:"sender_user_name"` // Name of the user who sent the message originally
}

func (m MessageOriginHiddenUser) MessageOriginType() string { return m.Type }
func (m MessageOriginHiddenUser) MessageOriginDate() int    { return m.Date }
