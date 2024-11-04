package types

// Giveaway https://core.telegram.org/bots/api#giveaway
type Giveaway struct {
	Chats []Chat `json:"chats"`
}
