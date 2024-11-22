package types

// MessageEntity https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	// Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag or #hashtag@chatusername), “cashtag” ($USD or $USD@chatusername), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “blockquote” (block quotation), “expandable_blockquote” (collapsed-by-default block quotation), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units
	Length int `json:"length"`
	// Optional: URL for "text_link" entities
	URL string `json:"url,omitempty"`
	// Optional: User for "text_mention" entities
	User *User `json:"user,omitempty"`
	// Optional: Programming language for "pre" entities
	Language string `json:"language,omitempty"`
	// Optional: For “custom_emoji” only, unique identifier of the custom emoji. To get full information about the sticker use method https://core.telegram.org/bots/api#getcustomemojistickers
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}
