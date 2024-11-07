package types

type BusinessConnection struct {
	ID         string `json:"id"`           // Unique identifier of the business connection
	User       User   `json:"user"`         // Business account user who created the connection
	UserChatID int64  `json:"user_chat_id"` // Identifier of a private chat with the user
	Date       int    `json:"date"`         // Date of connection establishment in Unix time
	CanReply   bool   `json:"can_reply"`    // True if the bot can act on behalf of the business account in recent chats
	IsEnabled  bool   `json:"is_enabled"`   // True if the connection is active
}
