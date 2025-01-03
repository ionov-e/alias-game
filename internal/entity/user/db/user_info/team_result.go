package user_info

type TeamInfo struct {
	Name         string        `json:"n"`  //nolint:tagliatelle
	RoundResults []RoundResult `json:"rr"` //nolint:tagliatelle
}

type RoundResult struct {
	CorrectAnswersCount   uint16
	IncorrectAnswersCount uint16
	SkippedAnswersCount   uint16
}
