package types

type ReactionTypeCustomEmoji struct {
	// Always “custom_emoji”
	Type string `json:"type"`
	// Custom emoji identifier
	CustomEmojiID string `json:"custom_emoji_id"`
}

func (r ReactionTypeCustomEmoji) ReactionType() string { return r.Type }
