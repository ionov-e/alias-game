package db

import (
	dictionaryConstant "alias-game/internal/constant/dictionary"
	menuConstant "alias-game/internal/constant/menu"
	userConstant "alias-game/internal/constant/user"
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
	PreferencePenaltyCost    float32                   `json:"ppc"`           //nolint:tagliatelle
	PreferenceWordDifficulty uint8                     `json:"pwd"`           //nolint:tagliatelle
	DictionaryHistory        []DictionaryCount         `json:"dc,omitempty"`  //nolint:tagliatelle
	RoundStartTime           time.Time                 `json:"rst"`           //nolint:tagliatelle
	RoundEndTime             time.Time                 `json:"ret"`           //nolint:tagliatelle
	RoundDictionaryKey       dictionaryConstant.Key    `json:"rdk"`           //nolint:tagliatelle
	RoundWords               []string                  `json:"rw,omitempty"`  //nolint:tagliatelle
	RoundWordResults         []userConstant.WordResult `json:"rwr,omitempty"` //nolint:tagliatelle
}

func (u *UserInfo) AddWordResult(wordNumber uint16, wordResult userConstant.WordResult) {
	u.AddLastRequest()
	u.CurrentMenu = string(menuConstant.NewWordKey(wordNumber))

	if u.RoundWordResults == nil {
		u.RoundWordResults = make([]userConstant.WordResult, wordNumber+1)
	}

	if int(wordNumber) >= len(u.RoundWordResults) {
		u.RoundWordResults = append(
			u.RoundWordResults,
			make(
				[]userConstant.WordResult,
				int(wordNumber)-len(u.RoundWordResults)+1,
			)...,
		)
	}

	u.RoundWordResults[wordNumber] = wordResult
}

func (u *UserInfo) AddLastRequest() {
	u.LastRequestTime = time.Now()
}

func (u *UserInfo) ChooseAnotherDictionary(dictionaryKey dictionaryConstant.Key) {
	u.RoundDictionaryKey = dictionaryKey

	if u.DictionaryHistory == nil {
		u.DictionaryHistory = []DictionaryCount{
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

	u.DictionaryHistory = append(u.DictionaryHistory, DictionaryCount{
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
		u.DictionaryHistory = []DictionaryCount{}
	}
	return nil
}
