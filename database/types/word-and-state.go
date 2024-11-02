package types

import (
	"encoding/json"
	"fmt"
)

type WordAndState struct {
	Word  string `json:"w"`
	State uint8  `json:"v"`
}

const (
	NotAnswered = iota
	Correct
	Incorrect
	Skipped
)

func (w *WordAndState) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(w)
	if err != nil {
		return nil, fmt.Errorf("marshal WordAndState failed: %w", err)
	}
	return data, nil
}

func (w *WordAndState) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &w)
	if err != nil {
		return fmt.Errorf("unmarshal WordAndState failed: %w", err)
	}
	return nil
}
