package types

// MessageEntity https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type          string  `json:"type"`                      // Type of the entity (e.g., "mention", "hashtag")
	Offset        int     `json:"offset"`                    // Start position in UTF-16 code units
	Length        int     `json:"length"`                    // Length of the entity in UTF-16 code units
	URL           *string `json:"url,omitempty"`             // Optional: URL for "text_link" entities
	User          *User   `json:"user,omitempty"`            // Optional: User for "text_mention" entities
	Language      *string `json:"language,omitempty"`        // Optional: Programming language for "pre" entities
	CustomEmojiID *string `json:"custom_emoji_id,omitempty"` // Optional: ID for "custom_emoji" entities
}
