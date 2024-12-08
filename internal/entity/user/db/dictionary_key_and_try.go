package db

import (
	dictionaryConstant "alias-game/internal/constant/dictionary"
	"strconv"
)

// DictionaryKeyAndTry represents a menu key stored in Redis.
type DictionaryKeyAndTry struct {
	BaseKey   dictionaryConstant.Key
	TryNumber uint16
}

func NewDictionaryKeyAndTry(baseKey dictionaryConstant.Key, tryNumber uint16) DictionaryKeyAndTry {
	return DictionaryKeyAndTry{
		BaseKey:   baseKey,
		TryNumber: tryNumber,
	}
}

func (d DictionaryKeyAndTry) String() string {
	return string(d.BaseKey) + ":" + strconv.Itoa(int(d.TryNumber))
}
