package types

type MessageOriginHiddenUser struct {
	// Type of the message origin, always "hidden_user"
	Type string `json:"type"`
	// Date the message was originally sent (Unix time)
	Date int `json:"date"`
	// Name of the user who sent the message originally
	SenderUserName string `json:"sender_user_name"`
}

func (m MessageOriginHiddenUser) MessageOriginType() string { return m.Type }
func (m MessageOriginHiddenUser) MessageOriginDate() int    { return m.Date }
