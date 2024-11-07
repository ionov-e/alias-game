package types

// ChatMemberBanned represents a chat member that was banned and can't return to the chat.
type ChatMemberBanned struct {
	Status    string `json:"status"`     // The member's status in the chat, always “kicked”.
	User      User   `json:"user"`       // Information about the user.
	UntilDate int64  `json:"until_date"` // Date when ban will be lifted, Unix time. If 0, user is banned forever.
}

// GetStatus returns the member's status.
func (c ChatMemberBanned) GetStatus() string {
	return c.Status
}

// GetUser returns information about the user.
func (c ChatMemberBanned) GetUser() User {
	return c.User
}
