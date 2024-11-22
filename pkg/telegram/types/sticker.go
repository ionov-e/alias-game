package types

// Sticker https://core.telegram.org/bots/api#sticker
type Sticker struct {
	// Identifier for this file, used to download or reuse it.
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
	Type string `json:"type"`
	// Sticker width.
	Width int `json:"width"`
	// Sticker height.
	Height int `json:"height"`
	// True if the sticker is animated.
	IsAnimated bool `json:"is_animated"`
	// True if the sticker is a video sticker.
	IsVideo bool `json:"is_video"`
	// Optional. Thumbnail of the sticker in .WEBP or .JPG format.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// Optional. Emoji associated with the sticker.
	Emoji string `json:"emoji,omitempty"`
	// Optional. Name of the sticker set to which the sticker belongs.
	SetName string `json:"set_name,omitempty"`
	// Optional. Premium animation for premium stickers.
	PremiumAnimation *File `json:"premium_animation,omitempty"`
	// Optional. Position where the mask should be placed for mask stickers.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// Optional. Unique identifier of the custom emoji for custom emoji stickers.
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
	// Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
	NeedsRepainting bool `json:"needs_repainting,omitempty"`
	// Optional. Size of the sticker file in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}
