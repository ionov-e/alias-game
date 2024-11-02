package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserInfo struct {
	TelegramID               int64     `json:"i"`
	Name                     string    `json:"n"`
	BlockedUntil             time.Time `json:"bu"`
	BlockedTimes             uint16    `json:"bt"`
	LastRequestTime          time.Time `json:"lrt"`
	FirstFrequentRequestTime time.Time `json:"ffrt"`
	// Same as in Telegram for user: IETF language tag
	PreferenceLanguage string `json:"pl"`
	// In seconds
	PreferenceRoundTime uint16 `json:"prt"`
	// Number of points to reduce for wrong answers
	PreferencePenaltyCost    float32        `json:"ppc"`
	PreferenceWordDifficulty uint8          `json:"pwd"`
	PreferenceWordTopics     []uint16       `json:"pwt"`
	RoundStartTime           time.Time      `json:"rst"`
	RoundEndTime             time.Time      `json:"ret"`
	RoundWords               []WordAndState `json:"rw,omitempty"`
}

func (u *UserInfo) ResultStringForTelegram() string {
	if u.RoundWords == nil {
		return ""
	}

	msg := "Round results:\n"
	// return every word in roundWords with its state
	for _, wordAndState := range u.RoundWords {
		switch wordAndState.State {
		case Correct:
			msg += wordAndState.Word + " ✅\n"
		case Incorrect:
			msg += wordAndState.Word + " ❌\n"
		case Skipped:
			msg += wordAndState.Word + " ❓\n"
		case NotAnswered:
			msg += wordAndState.Word + " ❔\n"
		}
	}
	return msg
}

func (u *UserInfo) AddNewWord(word string) {
	u.RoundWords = append(u.RoundWords, WordAndState{
		Word:  word,
		State: NotAnswered,
	})
}

func (u *UserInfo) LastWord() WordAndState {
	return u.RoundWords[len(u.RoundWords)-1]
}

func (u *UserInfo) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("marshal UserInfo failed: %w", err)
	}
	return data, nil
}

func (u *UserInfo) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &u)
	if err != nil {
		return fmt.Errorf("unmarshal UserInfo failed: %w", err)
	}
	if u.RoundWords == nil {
		u.RoundWords = []WordAndState{}
	}
	return nil
}
