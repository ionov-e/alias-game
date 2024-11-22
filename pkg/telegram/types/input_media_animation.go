package types

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	// Type of the result, must be "animation"
	Type string `json:"type"`
	// File to send as file_id, HTTP URL, or "attach://<file_attach_name>"
	Media string `json:"media"`
	// Optional: Thumbnail of the file in JPEG format
	Thumbnail *InputFileOrString `json:"thumbnail,omitempty"`
	// Optional: Caption of the animation
	Caption string `json:"caption,omitempty"`
	// Optional: Mode for parsing entities in the caption
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional: List of special entities in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional: Show caption above media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`
	// Optional: Animation width
	Width int `json:"width,omitempty"`
	// Optional: Animation height
	Height int `json:"height,omitempty"`
	// Optional: Animation duration in seconds
	Duration int `json:"duration,omitempty"`
	// Optional: Cover with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// InputMediaType returns the type of the media, which is "animation".
func (a *InputMediaAnimation) InputMediaType() string {
	return a.Type
}
