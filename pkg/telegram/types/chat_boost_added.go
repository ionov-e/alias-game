package types

// ChatBoostAdded
// This object represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	// Number of boosts added by the user
	BoostCount int `json:"boost_count"`
}
