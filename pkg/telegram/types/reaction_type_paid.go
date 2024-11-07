package types

type ReactionTypePaid struct {
	Type string `json:"type"` // Always “paid”
}

func (r ReactionTypePaid) ReactionType() string { return r.Type }
