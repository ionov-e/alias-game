package types

// ChatMemberMember represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	Status    string `json:"status"`               // The member's status in the chat, always “member”.
	User      User   `json:"user"`                 // Information about the user.
	UntilDate *int64 `json:"until_date,omitempty"` // Optional. Date when the user's subscription expires, Unix time.
}

// GetStatus returns the member's status.
func (c ChatMemberMember) GetStatus() string {
	return c.Status
}

// GetUser returns information about the user.
func (c ChatMemberMember) GetUser() User {
	return c.User
}
