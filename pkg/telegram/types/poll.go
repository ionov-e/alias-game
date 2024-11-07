package types

// Poll represents a poll in a message.
type Poll struct {
	ID                    string           `json:"id"`                             // Unique poll identifier.
	Question              string           `json:"question"`                       // Poll question.
	QuestionEntities      *[]MessageEntity `json:"question_entities,omitempty"`    // Optional. Special entities in the question.
	Options               []PollOption     `json:"options"`                        // List of poll options.
	TotalVoterCount       int              `json:"total_voter_count"`              // Total number of users who voted.
	IsClosed              bool             `json:"is_closed"`                      // True if the poll is closed.
	IsAnonymous           bool             `json:"is_anonymous"`                   // True if the poll is anonymous.
	Type                  string           `json:"type"`                           // Poll type, can be "regular" or "quiz".
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`        // True if multiple answers are allowed.
	CorrectOptionID       *int             `json:"correct_option_id,omitempty"`    // Optional. Correct answer for quiz-type polls.
	Explanation           *string          `json:"explanation,omitempty"`          // Optional. Text shown when a user answers incorrectly.
	ExplanationEntities   *[]MessageEntity `json:"explanation_entities,omitempty"` // Optional. Special entities in the explanation.
	OpenPeriod            *int             `json:"open_period,omitempty"`          // Optional. Duration the poll is active after creation.
	CloseDate             *int             `json:"close_date,omitempty"`           // Optional. Timestamp when the poll closes.
}
