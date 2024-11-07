package types

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`  // Chat which was boosted
	Boost ChatBoost `json:"boost"` // Information about the chat boost
}
