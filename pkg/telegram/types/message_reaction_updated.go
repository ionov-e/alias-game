package types

type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`                 // The chat containing the message the user reacted to
	MessageID   int            `json:"message_id"`           // Unique identifier of the message inside the chat
	User        *User          `json:"user,omitempty"`       // The user that changed the reaction, if the user isn't anonymous
	ActorChat   *Chat          `json:"actor_chat,omitempty"` // The chat on behalf of which the reaction was changed, if the user is anonymous
	Date        int            `json:"date"`                 // Date of the change in Unix time
	OldReaction []ReactionType `json:"old_reaction"`         // Previous list of reaction types that were set by the user
	NewReaction []ReactionType `json:"new_reaction"`         // New list of reaction types that have been set by the user
}
