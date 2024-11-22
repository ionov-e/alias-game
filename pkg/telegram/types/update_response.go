package types

type UpdateResponse struct {
	Ok          bool     `json:"ok"`
	Description string   `json:"description"` // Exists only if Ok is false
	Result      []Update `json:"result"`
}
