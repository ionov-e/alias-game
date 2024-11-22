package types

type ReactionCount struct {
	// Type of the reaction
	Type ReactionType `json:"type"`
	// Number of times the reaction was added
	TotalCount int `json:"total_count"`
}
