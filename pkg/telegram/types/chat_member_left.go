package types

// ChatMemberLeft represents a chat member who isn't currently a member of the chat.
type ChatMemberLeft struct {
	Status string `json:"status"` // The member's status in the chat, always “left”.
	User   User   `json:"user"`   // Information about the user.
}

// GetStatus returns the member's status.
func (c ChatMemberLeft) GetStatus() string {
	return c.Status
}

// GetUser returns information about the user.
func (c ChatMemberLeft) GetUser() User {
	return c.User
}
