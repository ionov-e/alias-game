package types

// PollAnswer represents an answer from a user in a non-anonymous poll. https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	// Unique poll identifier.
	PollID string `json:"poll_id"`
	// Optional. The chat that changed the answer.
	VoterChat *Chat `json:"voter_chat,omitempty"`
	// Optional. The user that changed the answer.
	User *User `json:"user,omitempty"`
	// 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
	OptionIDs []int `json:"option_ids"`
}
