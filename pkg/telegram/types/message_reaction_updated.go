package types

type MessageReactionUpdated struct {
	// The chat containing the message the user reacted to
	Chat Chat `json:"chat"`
	// Unique identifier of the message inside the chat
	MessageID int `json:"message_id"`
	// Optional. The user that changed the reaction, if the user isn't anonymous
	User *User `json:"user,omitempty"`
	// Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	ActorChat *Chat `json:"actor_chat,omitempty"`
	// Date of the change in Unix time
	Date int `json:"date"`
	// Previous list of reaction types that were set by the user
	OldReaction []ReactionType `json:"old_reaction"`
	// New list of reaction types that have been set by the user
	NewReaction []ReactionType `json:"new_reaction"`
}
