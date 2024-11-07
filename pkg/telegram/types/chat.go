package types

// Chat https://core.telegram.org/bots/api#chat represents a chat entity with basic information.
type Chat struct {
	ID        int64   `json:"id"`                   // Unique identifier for this chat. Safe to store as a signed 64-bit integer.
	Type      string  `json:"type"`                 // Type of the chat, can be "private", "group", "supergroup", or "channel".
	Title     *string `json:"title,omitempty"`      // Optional. Title for supergroups, channels, and group chats.
	Username  *string `json:"username,omitempty"`   // Optional. Username for private chats, supergroups, and channels if available.
	FirstName *string `json:"first_name,omitempty"` // Optional. First name of the other party in a private chat.
	LastName  *string `json:"last_name,omitempty"`  // Optional. Last name of the other party in a private chat.
	IsForum   *bool   `json:"is_forum,omitempty"`   // Optional. True if the supergroup chat is a forum (has topics enabled).
}
