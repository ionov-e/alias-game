package types

// MessageResponse https://core.telegram.org/bots/api#sendmessage
type MessageResponse struct {
	Ok          bool    `json:"ok"`
	Description *string `json:"description"` // Exists only if Ok is false
	Result      Message `json:"result"`
}

func (m *MessageResponse) IsOk() bool {
	return m.Ok
}

func (m *MessageResponse) DescriptionText() string {
	if m.Description != nil {
		return *m.Description
	}
	return "For some reason it is empty"
}
