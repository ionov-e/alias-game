package types

type InputSticker struct {
	Sticker      interface{}   `json:"sticker"`                 // InputFile or String: The sticker file to add
	Format       string        `json:"format"`                  // Format of the sticker, must be "static", "animated", or "video"
	EmojiList    []string      `json:"emoji_list"`              // List of 1-20 emojis associated with the sticker
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional: Position for mask stickers
	Keywords     *[]string     `json:"keywords,omitempty"`      // Optional: List of up to 20 keywords for search
}
