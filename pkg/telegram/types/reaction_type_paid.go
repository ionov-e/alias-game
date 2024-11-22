package types

type ReactionTypePaid struct {
	// Always “paid”
	Type string `json:"type"`
}

func (r ReactionTypePaid) ReactionType() string { return r.Type }
