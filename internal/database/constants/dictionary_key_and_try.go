package constants

import "strconv"

// DictionaryKeyAndTry represents a menu key stored in Redis.
type DictionaryKeyAndTry struct {
	BaseKey   DictionaryKey
	TryNumber uint16
}

func NewDictionaryKey(baseKey DictionaryKey, index uint16) DictionaryKeyAndTry {
	return DictionaryKeyAndTry{
		BaseKey:   baseKey,
		TryNumber: index,
	}
}

func (d DictionaryKeyAndTry) String() string {
	return string(d.BaseKey) + ":" + strconv.Itoa(int(d.TryNumber))
}
