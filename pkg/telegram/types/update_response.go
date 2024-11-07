package types

type UpdateResponse struct {
	Ok          bool     `json:"ok"`
	Description *string  `json:"description"` // Exists only if Ok is false
	Result      []Update `json:"result"`
}

func (u *UpdateResponse) DescriptionText() string {
	if u.Description != nil {
		return *u.Description
	}
	return "For some reason it is empty"
}
