package types

// Game https://core.telegram.org/bots/api#game
type Game struct {
	Title        string          `json:"title"`                   // Title of the game.
	Description  string          `json:"description"`             // Description of the game.
	Photo        []PhotoSize     `json:"photo"`                   // Array of photos to be displayed in the game message in chats.
	Text         *string         `json:"text,omitempty"`          // Optional. Brief description or high scores included in the game message.
	TextEntities []MessageEntity `json:"text_entities,omitempty"` // Optional. Special entities in text, such as usernames, URLs, and bot commands.
	Animation    *Animation      `json:"animation,omitempty"`     // Optional. Animation displayed in the game message in chats, uploaded via BotFather.
}
