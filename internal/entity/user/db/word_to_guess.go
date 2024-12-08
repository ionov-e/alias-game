package db

import (
	userConstant "alias-game/internal/constant/user"
	"encoding/json"
	"fmt"
)

type WordToGuess struct {
	NumberInDictionary uint16                  `json:"n"` //nolint:tagliatelle
	Result             userConstant.WordResult `json:"s"` //nolint:tagliatelle
}

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
