package user_info

import (
	dictionaryConstant "alias-game/internal/constant/dictionary"
)

type DictionaryCount struct {
	DictionaryKey dictionaryConstant.Key `json:"d"` //nolint:tagliatelle
	Count         uint16                 `json:"c"` //nolint:tagliatelle
}
