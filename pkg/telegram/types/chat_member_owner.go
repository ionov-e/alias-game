package types

// ChatMemberOwner represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	Status      string `json:"status"`                 // The member's status in the chat, always “creator”.
	User        User   `json:"user"`                   // Information about the user.
	IsAnonymous bool   `json:"is_anonymous"`           // True if the user's presence in the chat is hidden.
	CustomTitle string `json:"custom_title,omitempty"` // Optional. Custom title for this user.
}

// GetStatus returns the member's status.
func (c ChatMemberOwner) GetStatus() string {
	return c.Status
}

// GetUser returns information about the user.
func (c ChatMemberOwner) GetUser() User {
	return c.User
}
