package vo

import (
	dictionaryConstant "alias-game/internal/constant/dictionary"
	userConstant "alias-game/internal/constant/user"
	"encoding/json"
	"fmt"
	"time"
)

// UserInfo represents Value object stored
type UserInfo struct {
	TelegramID               int64     `json:"i"`    //nolint:tagliatelle
	CurrentMenu              string    `json:"cm"`   //nolint:tagliatelle
	Name                     string    `json:"n"`    //nolint:tagliatelle
	BlockedUntil             time.Time `json:"bu"`   //nolint:tagliatelle
	BlockedTimes             uint16    `json:"bt"`   //nolint:tagliatelle
	LastRequestTime          time.Time `json:"lrt"`  //nolint:tagliatelle
	FirstFrequentRequestTime time.Time `json:"ffrt"` //nolint:tagliatelle
	// Same as in Telegram for user: IETF language tag
	PreferenceLanguage string `json:"pl"` //nolint:tagliatelle
	// In seconds
	PreferenceRoundTime uint16 `json:"prt"` //nolint:tagliatelle
	// Number of points to reduce for wrong answers
	PreferencePenaltyCost    float32                `json:"ppc"`          //nolint:tagliatelle
	PreferenceWordDifficulty uint8                  `json:"pwd"`          //nolint:tagliatelle
	DictionaryHistory        []dictionaryCount      `json:"dc,omitempty"` //nolint:tagliatelle
	RoundStartTime           time.Time              `json:"rst"`          //nolint:tagliatelle
	RoundEndTime             time.Time              `json:"ret"`          //nolint:tagliatelle
	RoundDictionaryKey       dictionaryConstant.Key `json:"rdk"`          //nolint:tagliatelle
	// Starts with 0 (index in RoundWords slice)
	RoundWordNumber uint16 `json:"rwn"` //nolint:tagliatelle
	// Starts with 0 (index in AllTeamsInfo slice)
	RoundTeamNumber  uint16                    `json:"rtn"`           //nolint:tagliatelle
	RoundWords       []string                  `json:"rw,omitempty"`  //nolint:tagliatelle
	RoundWordResults []userConstant.WordResult `json:"rwr,omitempty"` //nolint:tagliatelle
	WordCountToWin   uint16                    `json:"wctw"`          //nolint:tagliatelle
	AllTeamsInfo     []TeamInfo                `json:"ati,omitempty"` //nolint:tagliatelle
}

type TeamInfo struct {
	Name         string        `json:"n"`  //nolint:tagliatelle
	RoundResults []roundResult `json:"rr"` //nolint:tagliatelle
}

type roundResult struct {
	CorrectAnswersCount   uint16
	IncorrectAnswersCount uint16
	SkippedAnswersCount   uint16
}

type dictionaryCount struct {
	DictionaryKey dictionaryConstant.Key `json:"d"` //nolint:tagliatelle
	Count         uint16                 `json:"c"` //nolint:tagliatelle
}

func (u *UserInfo) Word(wordNumber uint16) (string, error) {
	if int(wordNumber) >= len(u.RoundWords)-1 {
		return "", fmt.Errorf("wordNumber %d is too much for RoundWords slice of %d", wordNumber, len(u.RoundWords))
	}
	return u.RoundWords[wordNumber], nil
}

func (u *UserInfo) SetRoundWordResult(wordNumber uint16, wordResult userConstant.WordResult) {
	u.AddLastRequest()

	if u.RoundWordResults == nil {
		u.RoundWordResults = []userConstant.WordResult{}
	}

	if int(wordNumber) >= len(u.RoundWordResults) {
		u.RoundWordResults = append(
			u.RoundWordResults,
			wordResult,
		)
	} else {
		u.RoundWordResults[wordNumber] = wordResult
	}
}

func (u *UserInfo) AddRoundResult(correctAnswers, incorrectAnswers, skippedAnswers int) error {
	err := u.checkAllTeamsInfoComparedToRoundNumber()
	if err != nil {
		return fmt.Errorf("failed check in AddRoundResult: %w", err)
	}

	currentTeam := &u.AllTeamsInfo[u.RoundTeamNumber]

	currentTeam.RoundResults = append(
		currentTeam.RoundResults,
		roundResult{
			CorrectAnswersCount:   uint16(correctAnswers),
			IncorrectAnswersCount: uint16(incorrectAnswers),
			SkippedAnswersCount:   uint16(skippedAnswers),
		},
	)

	return nil
}

func (u *UserInfo) checkAllTeamsInfoComparedToRoundNumber() error {
	if u.AllTeamsInfo == nil {
		return fmt.Errorf("allTeamsInfo is nil, but RoundTeamNumber is %d", u.RoundTeamNumber)
	}

	allTeamCount := u.AllTeamCount()

	if allTeamCount == 1 {
		return nil
	}

	if u.RoundTeamNumber >= allTeamCount {
		return fmt.Errorf("RoundTeamNumber is %d, but AllTeamsInfo count is %d", u.RoundTeamNumber, allTeamCount)
	}
	return nil
}

func (u *UserInfo) IsGameEnd() (bool, error) {
	err := u.checkAllTeamsInfoComparedToRoundNumber()
	if err != nil {
		return false, fmt.Errorf("failed check in IsGameEnded: %w", err)
	}

	// All teams should have equal number of rounds before check
	maxRounds := len(u.AllTeamsInfo[0].RoundResults)
	for _, team := range u.AllTeamsInfo[1:] {
		if len(team.RoundResults) != maxRounds {
			return false, nil
		}
	}

	for _, team := range u.AllTeamsInfo {
		totalCorrectForTeam := uint16(0)
		for _, round := range team.RoundResults {
			totalCorrectForTeam += round.CorrectAnswersCount
		}

		if totalCorrectForTeam >= u.WordCountToWin {
			return true, nil
		}
	}

	return false, nil
}

func (u *UserInfo) AddLastRequest() {
	u.LastRequestTime = time.Now()
}

func (u *UserInfo) ChooseAnotherDictionary(dictionaryKey dictionaryConstant.Key) {
	u.RoundDictionaryKey = dictionaryKey

	if u.DictionaryHistory == nil {
		u.DictionaryHistory = []dictionaryCount{
			{DictionaryKey: dictionaryKey, Count: 1},
		}
		return
	}

	for i, dictCount := range u.DictionaryHistory {
		if dictCount.DictionaryKey == dictionaryKey {
			u.DictionaryHistory[i].Count++
			return
		}
	}

	u.DictionaryHistory = append(u.DictionaryHistory, dictionaryCount{
		DictionaryKey: dictionaryKey,
		Count:         1,
	})
}

func (u UserInfo) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("marshal userInfo failed: %w", err)
	}
	return data, nil
}

func (u *UserInfo) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &u)
	if err != nil {
		return fmt.Errorf("unmarshal userInfo failed: %w", err)
	}
	if u.RoundWordResults == nil {
		u.RoundWordResults = []userConstant.WordResult{}
	}
	if u.RoundWords == nil {
		u.RoundWords = []string{}
	}
	if u.DictionaryHistory == nil {
		u.DictionaryHistory = []dictionaryCount{}
	}
	return nil
}

func (u UserInfo) AllTeamCount() uint16 {
	return uint16(len(u.AllTeamsInfo))
}
