package types

// ForumTopicEdited
// This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	// Name New name of the topic, if it was edited
	Name string `json:"name,omitempty"`
	// IconCustomEmojiId New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}
