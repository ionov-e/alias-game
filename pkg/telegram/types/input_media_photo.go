package types

// InputMediaPhoto represents a photo to be sent.
type InputMediaPhoto struct {
	Type                  string           `json:"type"`                               // Type of the result, must be "photo"
	Media                 string           `json:"media"`                              // File to send as file_id, HTTP URL, or "attach://<file_attach_name>"
	Caption               *string          `json:"caption,omitempty"`                  // Optional: Caption of the photo
	ParseMode             *string          `json:"parse_mode,omitempty"`               // Optional: Mode for parsing entities in the caption
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`         // Optional: List of special entities in the caption
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"` // Optional: Show caption above media
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`              // Optional: Cover with a spoiler animation
}

// InputMediaType returns the type of the media, which is "photo".
func (p *InputMediaPhoto) InputMediaType() string {
	return p.Type
}
