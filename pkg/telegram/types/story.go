package types

// Story https://core.telegram.org/bots/api#story represents a story posted in a chat.
type Story struct {
	// Chat that posted the story.
	Chat Chat `json:"chat"`
	// Unique identifier for the story in the chat.
	ID int64 `json:"id"`
}
