package types

// PollAnswer represents an answer from a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string `json:"poll_id"`              // Unique poll identifier.
	VoterChat *Chat  `json:"voter_chat,omitempty"` // Optional. The chat that changed the answer.
	User      *User  `json:"user,omitempty"`       // Optional. The user that changed the answer.
	OptionIDs []int  `json:"option_ids"`           // 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
}
