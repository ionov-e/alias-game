package constants

import (
	"fmt"
	"strconv"
	"strings"
)

const wordPrefix string = "w-" // Dynamic keys, e.g., "w-123"

// MenuKeyStored represents a menu key stored in Redis.
type MenuKeyStored string

const (
	Empty                    MenuKeyStored = ""
	MenuStart0Key            MenuKeyStored = "st-0"
	MenuDictionaryChoice0Key MenuKeyStored = "di-ch-0"
)

func NewWordMenuKey(number uint16) MenuKeyStored {
	return MenuKeyStored(fmt.Sprintf("%s%d", wordPrefix, number))
}

func (mk MenuKeyStored) String() string {
	return string(mk)
}

// IsWord checks if the MenuKeyStored has the wordPrefix.
func (mk MenuKeyStored) IsWord() bool {
	return strings.HasPrefix(string(mk), wordPrefix)
}

func (mk MenuKeyStored) WordNumber() (uint16, error) {
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
