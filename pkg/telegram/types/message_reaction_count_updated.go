package types

type MessageReactionCountUpdated struct {
	// The chat containing the message
	Chat Chat `json:"chat"`
	// Unique message identifier inside the chat
	MessageID int `json:"message_id"`
	// Date of the change in Unix time
	Date int `json:"date"`
	// List of reactions that are present on the message
	Reactions []ReactionCount `json:"reactions"`
}
