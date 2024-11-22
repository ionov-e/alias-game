package types

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`                  // Chat to which the request was sent
	From       User            `json:"from"`                  // User that sent the join request
	UserChatID int64           `json:"user_chat_id"`          // Identifier of a private chat with the user
	Date       int             `json:"date"`                  // Date the request was sent in Unix time
	Bio        string          `json:"bio,omitempty"`         // Optional. Bio of the user
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"` // Optional. Chat invite link that was used by the user to send the join request
}
