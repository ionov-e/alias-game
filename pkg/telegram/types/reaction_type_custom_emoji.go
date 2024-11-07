package types

type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`            // Always “custom_emoji”
	CustomEmojiID string `json:"custom_emoji_id"` // Custom emoji identifier
}

func (r ReactionTypeCustomEmoji) ReactionType() string { return r.Type }
