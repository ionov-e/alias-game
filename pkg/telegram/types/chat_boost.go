package types

type ChatBoost struct {
	BoostID        string          `json:"boost_id"`        // Unique identifier of the boost
	AddDate        int             `json:"add_date"`        // Point in time (Unix timestamp) when the chat was boosted
	ExpirationDate int             `json:"expiration_date"` // Point in time (Unix timestamp) when the boost will expire
	Source         ChatBoostSource `json:"source"`          // Source of the added boost
}
