package types

type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`        // Chat which was boosted
	BoostID    string          `json:"boost_id"`    // Unique identifier of the boost
	RemoveDate int             `json:"remove_date"` // Point in time (Unix timestamp) when the boost was removed
	Source     ChatBoostSource `json:"source"`      // Source of the removed boost
}
