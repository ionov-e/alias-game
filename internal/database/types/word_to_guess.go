package types

import (
	"encoding/json"
	"fmt"
)

type WordToGuess struct {
	Word  string `json:"w"`
	State uint8  `json:"v"`
}

const (
	NotAnswered = iota
	Correct
	Incorrect
	Skipped
)

func (w WordToGuess) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(w)
	if err != nil {
		return nil, fmt.Errorf("marshal WordToGuess failed: %w", err)
	}
	return data, nil
}

func (w *WordToGuess) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &w)
	if err != nil {
		return fmt.Errorf("unmarshal WordToGuess failed: %w", err)
	}
	return nil
}
