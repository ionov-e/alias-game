package types

// Chat https://core.telegram.org/bots/api#chat
type Chat struct {
	// Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it
	ID int64 `json:"id"`
	// Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// Optional: for private chats, supergroups and channels if available
	Username string `json:"username"`
}
