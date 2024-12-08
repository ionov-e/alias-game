package menu

import (
	"fmt"
	"strconv"
	"strings"
)

const wordPrefix string = "w-" // Dynamic keys, e.g., "w-123"

// Key represents a menu key stored
type Key string

const (
	Start0Key            Key = "st-0"
	DictionaryChoice0Key Key = "di-ch-0"
)

func NewWordKey(number uint16) Key {
	return Key(fmt.Sprintf("%s%d", wordPrefix, number))
}

// IsWord checks if the Key has the wordPrefix.
func (mk Key) IsWord() bool {
	return strings.HasPrefix(string(mk), wordPrefix)
}

func (mk Key) WordNumber() (uint16, error) {
	parts := strings.SplitN(string(mk), "-", 2)
	if len(parts) != 2 {
		return 0, fmt.Errorf("key %s has no number after wordPrefix", mk)
	}
	number, err := strconv.ParseUint(parts[1], 10, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid number in key %s: %w", mk, err)
	}
	return uint16(number), nil
}
