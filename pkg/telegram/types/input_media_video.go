package types

// InputMediaVideo represents a video to be sent.
type InputMediaVideo struct {
	Type                  string           `json:"type"`                               // Type of the result, must be "video"
	Media                 string           `json:"media"`                              // File to send as file_id, HTTP URL, or "attach://<file_attach_name>"
	Thumbnail             *interface{}     `json:"thumbnail,omitempty"`                // Optional: Thumbnail of the file in JPEG format
	Caption               *string          `json:"caption,omitempty"`                  // Optional: Caption of the video
	ParseMode             *string          `json:"parse_mode,omitempty"`               // Optional: Mode for parsing entities in the caption
	CaptionEntities       *[]MessageEntity `json:"caption_entities,omitempty"`         // Optional: List of special entities in the caption
	ShowCaptionAboveMedia *bool            `json:"show_caption_above_media,omitempty"` // Optional: Show caption above media
	Width                 *int             `json:"width,omitempty"`                    // Optional: Video width
	Height                *int             `json:"height,omitempty"`                   // Optional: Video height
	Duration              *int             `json:"duration,omitempty"`                 // Optional: Video duration in seconds
	SupportsStreaming     *bool            `json:"supports_streaming,omitempty"`       // Optional: Suitable for streaming
	HasSpoiler            *bool            `json:"has_spoiler,omitempty"`              // Optional: Cover with a spoiler animation
}

// InputMediaType returns the type of the media, which is "video".
func (v *InputMediaVideo) InputMediaType() string {
	return v.Type
}
