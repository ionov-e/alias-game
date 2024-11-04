package types

// MessageResponse https://core.telegram.org/bots/api#sendmessage
type MessageResponse struct {
	Ok          bool    `json:"ok"`
	Description string  `json:"description"`
	Result      Message `json:"result"`
}

func (m *MessageResponse) IsOk() bool {
	return m.Ok
}
