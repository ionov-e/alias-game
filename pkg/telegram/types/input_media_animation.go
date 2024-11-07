package types

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Type                  string           `json:"type"`                               // Type of the result, must be "animation"
	Media                 string           `json:"media"`                              // File to send as file_id, HTTP URL, or "attach://<file_attach_name>"
	Thumbnail             *interface{}     `json:"thumbnail,omitempty"`                // Optional: Thumbnail of the file in JPEG format
	Caption               *string          `json:"caption,omitempty"`                  // Optional: Caption of the animation
	ParseMode             *string          `json:"parse_mode,omitempty"`               // Optional: Mode for parsing entities in the caption
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`         // Optional: List of special entities in the caption
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"` // Optional: Show caption above media
	Width                 *int             `json:"width,omitempty"`                    // Optional: Animation width
	Height                *int             `json:"height,omitempty"`                   // Optional: Animation height
	Duration              *int             `json:"duration,omitempty"`                 // Optional: Animation duration in seconds
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`              // Optional: Cover with a spoiler animation
}

// InputMediaType returns the type of the media, which is "animation".
func (a *InputMediaAnimation) InputMediaType() string {
	return a.Type
}
