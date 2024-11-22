package types

// Poll represents a poll in a message.
type Poll struct {
	// Unique poll identifier.
	ID string `json:"id"`
	// Poll question.
	Question string `json:"question"`
	// Optional. Special entities in the question.
	QuestionEntities []MessageEntity `json:"question_entities,omitempty"`
	// List of poll options.
	Options []PollOption `json:"options"`
	// Total number of users who voted.
	TotalVoterCount int `json:"total_voter_count"`
	// True if the poll is closed.
	IsClosed bool `json:"is_closed"`
	// True if the poll is anonymous.
	IsAnonymous bool `json:"is_anonymous"`
	// Poll type, can be "regular" or "quiz".
	Type string `json:"type"`
	// True if multiple answers are allowed.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bo
	CorrectOptionID int `json:"correct_option_id,omitempty"`
	// Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation,omitempty"`
	// Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
	// 	Optional. Amount of time in seconds the poll will be active after creation
	OpenPeriod int `json:"open_period,omitempty"`
	// Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate int `json:"close_date,omitempty"`
}
