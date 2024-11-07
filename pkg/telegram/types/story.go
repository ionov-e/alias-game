package types

// Story https://core.telegram.org/bots/api#story represents a story posted in a chat.
type Story struct {
	Chat Chat  `json:"chat"` // Chat that posted the story.
	ID   int64 `json:"id"`   // Unique identifier for the story in the chat.
}
