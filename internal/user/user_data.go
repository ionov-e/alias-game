package user

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// data is a stored value object that contains data about User
type data struct {
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
	PreferencePenaltyCost    float32           `json:"ppc"`          //nolint:tagliatelle
	PreferenceWordDifficulty uint8             `json:"pwd"`          //nolint:tagliatelle
	DictionaryHistory        []dictionaryCount `json:"dc,omitempty"` //nolint:tagliatelle
	RoundStartTime           time.Time         `json:"rst"`          //nolint:tagliatelle
	RoundEndTime             time.Time         `json:"ret"`          //nolint:tagliatelle
	RoundDictionaryKey       DictionaryKey     `json:"rdk"`          //nolint:tagliatelle
	// Starts with 0 (index in RoundWords slice)
	RoundWordNumber uint16 `json:"rwn"` //nolint:tagliatelle
	// Starts with 0 (index in AllTeamsInfo slice)
	RoundTeamNumber  uint16       `json:"rtn"`           //nolint:tagliatelle
	RoundWords       []string     `json:"rw,omitempty"`  //nolint:tagliatelle
	RoundWordResults []WordResult `json:"rwr,omitempty"` //nolint:tagliatelle
	WordCountToWin   uint16       `json:"wctw"`          //nolint:tagliatelle
	AllTeamsInfo     []teamInfo   `json:"ati,omitempty"` //nolint:tagliatelle
}

type teamInfo struct {
	Name         string        `json:"n"`  //nolint:tagliatelle
	RoundResults []roundResult `json:"rr"` //nolint:tagliatelle
}

type roundResult struct {
	CorrectAnswersCount   uint16
	IncorrectAnswersCount uint16
	SkippedAnswersCount   uint16
}

type dictionaryCount struct {
	DictionaryKey DictionaryKey `json:"d"` //nolint:tagliatelle
	Count         uint16        `json:"c"` //nolint:tagliatelle
}

func (d *data) word(wordNumber uint16) (string, error) {
	if int(wordNumber) >= len(d.RoundWords)-1 {
		return "", fmt.Errorf("wordNumber %d is too much for RoundWords slice of %d", wordNumber, len(d.RoundWords))
	}
	return d.RoundWords[wordNumber], nil
}

func (d *data) setRoundWordResult(wordNumber uint16, wordResult WordResult) {
	d.addLastRequest()

	if d.RoundWordResults == nil {
		d.RoundWordResults = []WordResult{}
	}

	if int(wordNumber) >= len(d.RoundWordResults) {
		d.RoundWordResults = append(
			d.RoundWordResults,
			wordResult,
		)
	} else {
		d.RoundWordResults[wordNumber] = wordResult
	}
}

func (d *data) addRoundResult(correctAnswers, incorrectAnswers, skippedAnswers int) error {
	err := d.checkAllTeamsInfoComparedToRoundNumber()
	if err != nil {
		return fmt.Errorf("failed check in addRoundResult: %w", err)
	}

	currentTeam := &d.AllTeamsInfo[d.RoundTeamNumber]

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

func (d *data) checkAllTeamsInfoComparedToRoundNumber() error {
	if d.AllTeamsInfo == nil {
		return fmt.Errorf("allTeamsInfo is nil, but RoundTeamNumber is %d", d.RoundTeamNumber)
	}

	allTeamCount := d.allTeamCount()

	if allTeamCount == 1 {
		return nil
	}

	if d.RoundTeamNumber >= allTeamCount {
		return fmt.Errorf("RoundTeamNumber is %d, but AllTeamsInfo count is %d", d.RoundTeamNumber, allTeamCount)
	}
	return nil
}

func (d *data) isGameEnd() (bool, error) {
	err := d.checkAllTeamsInfoComparedToRoundNumber()
	if err != nil {
		return false, fmt.Errorf("failed check in IsGameEnded: %w", err)
	}

	// All teams should have equal number of rounds before check
	maxRounds := len(d.AllTeamsInfo[0].RoundResults)
	for _, team := range d.AllTeamsInfo[1:] {
		if len(team.RoundResults) != maxRounds {
			return false, nil
		}
	}

	for _, team := range d.AllTeamsInfo {
		totalCorrectForTeam := uint16(0)
		for _, round := range team.RoundResults {
			totalCorrectForTeam += round.CorrectAnswersCount
		}

		if totalCorrectForTeam >= d.WordCountToWin {
			return true, nil
		}
	}

	return false, nil
}

func (d *data) addLastRequest() {
	d.LastRequestTime = time.Now()
}

func (d *data) chooseAnotherDictionary(dictionaryKey DictionaryKey) {
	d.RoundDictionaryKey = dictionaryKey

	if d.DictionaryHistory == nil {
		d.DictionaryHistory = []dictionaryCount{
			{DictionaryKey: dictionaryKey, Count: 1},
		}
		return
	}

	for i, dictCount := range d.DictionaryHistory {
		if dictCount.DictionaryKey == dictionaryKey {
			d.DictionaryHistory[i].Count++
			return
		}
	}

	d.DictionaryHistory = append(d.DictionaryHistory, dictionaryCount{
		DictionaryKey: dictionaryKey,
		Count:         1,
	})
}

func (d *data) prepareRoundWordsFromDictionary() error {
	wordsFromDict, err := d.wordsFromDictionary()
	if err != nil {
		return fmt.Errorf("wordsFromDictionary failed: %w", err)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // We shouldn't bother
	r.Shuffle(len(wordsFromDict), func(i, j int) { wordsFromDict[i], wordsFromDict[j] = wordsFromDict[j], wordsFromDict[i] })
	d.RoundWords = wordsFromDict
	d.RoundWordResults = []WordResult{}
	d.RoundWordNumber = 0
	return nil
}

func (d data) wordsFromDictionary() ([]string, error) {
	if d.RoundDictionaryKey == Easy1 {
		return ease1List(), nil
	}

	return nil, fmt.Errorf("unknown dictionary key: %s", d.RoundDictionaryKey)
}

func (d data) convertTeamInfo() []TeamInfo {
	converted := make([]TeamInfo, len(d.AllTeamsInfo))

	for i, teamInfo := range d.AllTeamsInfo {
		roundResults := make([]RoundResult, len(teamInfo.RoundResults))
		var totalCorrect, totalIncorrect, totalSkipped uint16

		for j, rr := range teamInfo.RoundResults {
			roundResults[j] = RoundResult{
				CorrectAnswersCount:   rr.CorrectAnswersCount,
				IncorrectAnswersCount: rr.IncorrectAnswersCount,
				SkippedAnswersCount:   rr.SkippedAnswersCount,
			}
			totalCorrect += rr.CorrectAnswersCount
			totalIncorrect += rr.IncorrectAnswersCount
			totalSkipped += rr.SkippedAnswersCount
		}

		converted[i] = TeamInfo{
			Name:                       teamInfo.Name,
			RoundResults:               roundResults,
			TotalCorrectAnswersCount:   totalCorrect,
			TotalIncorrectAnswersCount: totalIncorrect,
			TotalSkippedAnswersCount:   totalSkipped,
		}
	}
	return converted
}

func (d data) allTeamCount() uint16 {
	return uint16(len(d.AllTeamsInfo))
}

func (d data) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("marshal data failed: %w", err)
	}
	return data, nil
}

func (d *data) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &d)
	if err != nil {
		return fmt.Errorf("unmarshal data failed: %w", err)
	}
	if d.RoundWordResults == nil {
		d.RoundWordResults = []WordResult{}
	}
	if d.RoundWords == nil {
		d.RoundWords = []string{}
	}
	if d.DictionaryHistory == nil {
		d.DictionaryHistory = []dictionaryCount{}
	}
	return nil
}
