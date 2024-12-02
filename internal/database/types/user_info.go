package types

import (
	dbConstants "alias-game/internal/database/constants"
	"encoding/json"
	"fmt"
	"time"
)

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
	PreferencePenaltyCost    float32                         `json:"ppc"`          //nolint:tagliatelle
	PreferenceWordDifficulty uint8                           `json:"pwd"`          //nolint:tagliatelle
	DictionaryHistory        []DictionaryCount               `json:"dc,omitempty"` //nolint:tagliatelle
	RoundStartTime           time.Time                       `json:"rst"`          //nolint:tagliatelle
	RoundEndTime             time.Time                       `json:"ret"`          //nolint:tagliatelle
	RoundDictionaryKeyAndTry dbConstants.DictionaryKeyAndTry `json:"rdk"`          //nolint:tagliatelle
	RoundWords               []WordToGuess                   `json:"rw,omitempty"` //nolint:tagliatelle
}

func (u *UserInfo) AddWordResult(wordNumber uint16, wordResult dbConstants.WordResult) { //TODO delete
	for i, word := range u.RoundWords {
		if word.NumberInDictionary == wordNumber {
			u.RoundWords[i].Result = wordResult
			return
		}
	}
	// If the word doesn't exist:
	u.RoundWords = append(u.RoundWords, WordToGuess{
		NumberInDictionary: wordNumber,
		Result:             wordResult,
	})
	u.CurrentMenu = string(dbConstants.NewWordMenuKey(wordNumber))
	u.AddLastRequest()
}

func (u *UserInfo) AddLastRequest() {
	u.LastRequestTime = time.Now()
}

func (u *UserInfo) FindDictionaryCountInHistory(dictionaryKey dbConstants.DictionaryKey) (*DictionaryCount, error) {
	for _, dictionaryCount := range u.DictionaryHistory {
		if dictionaryCount.DictionaryKey == dictionaryKey {
			return &dictionaryCount, nil
		}
	}
	return &DictionaryCount{}, fmt.Errorf("dictionaryKey key %s not found in history", dictionaryKey)
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
	if u.RoundWords == nil {
		u.RoundWords = []WordToGuess{}
	}
	return nil
}
