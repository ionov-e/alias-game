package types

// Sticker https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID           string        `json:"file_id"`                     // Identifier for this file, used to download or reuse it.
	FileUniqueID     string        `json:"file_unique_id"`              // Unique identifier for this file, constant over time and across bots.
	Type             string        `json:"type"`                        // Type of the sticker: "regular", "mask", or "custom_emoji".
	Width            int           `json:"width"`                       // Sticker width.
	Height           int           `json:"height"`                      // Sticker height.
	IsAnimated       bool          `json:"is_animated"`                 // True if the sticker is animated.
	IsVideo          bool          `json:"is_video"`                    // True if the sticker is a video sticker.
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`         // Optional. Thumbnail of the sticker in .WEBP or .JPG format.
	Emoji            *string       `json:"emoji,omitempty"`             // Optional. Emoji associated with the sticker.
	SetName          *string       `json:"set_name,omitempty"`          // Optional. Name of the sticker set to which the sticker belongs.
	PremiumAnimation *File         `json:"premium_animation,omitempty"` // Optional. Premium animation for premium stickers.
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`     // Optional. Position where the mask should be placed for mask stickers.
	CustomEmojiID    *string       `json:"custom_emoji_id,omitempty"`   // Optional. Unique identifier of the custom emoji for custom emoji stickers.
	NeedsRepainting  *bool         `json:"needs_repainting,omitempty"`  // Optional. True if the sticker requires repainting based on the context.
	FileSize         *int64        `json:"file_size,omitempty"`         // Optional. Size of the sticker file in bytes.
}
