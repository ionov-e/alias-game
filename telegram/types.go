package telegram

// Message https://core.telegram.org/bots/api#message
type Message struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
	} `json:"result"`
}
