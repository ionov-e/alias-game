package types

import dbConstants "alias-game/internal/database/constants"

type DictionaryCount struct {
	DictionaryKey dbConstants.DictionaryKey `json:"dictionary_prefix"`
	Count         uint16                    `json:"count"`
}
