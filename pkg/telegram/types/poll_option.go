package types

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	// Option text, 1-100 characters.
	Text string `json:"text"`
	// Optional. Special entities in the option text. Currently, only custom emoji entities are allowed in poll option texts
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
	// Number of users who voted for this option.
	VoterCount int `json:"voter_count"`
}
