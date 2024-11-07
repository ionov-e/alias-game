package types

// ForumTopicCreated
// This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	// Name of the topic
	Name string `json:"name"`
	// IconColor Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`
	// IconCustomEmojiId Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}
