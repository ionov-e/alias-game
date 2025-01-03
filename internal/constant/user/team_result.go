package user

type TeamInfo struct {
	Name                       string
	RoundResults               []RoundResult
	TotalCorrectAnswersCount   uint16
	TotalIncorrectAnswersCount uint16
	TotalSkippedAnswersCount   uint16
}

type RoundResult struct {
	CorrectAnswersCount   uint16
	IncorrectAnswersCount uint16
	SkippedAnswersCount   uint16
}
