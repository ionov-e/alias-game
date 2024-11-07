package types

type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`       // The chat containing the message
	MessageID int             `json:"message_id"` // Unique message identifier inside the chat
	Date      int             `json:"date"`       // Date of the change in Unix time
	Reactions []ReactionCount `json:"reactions"`  // List of reactions that are present on the message
}
