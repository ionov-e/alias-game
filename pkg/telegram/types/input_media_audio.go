package types

// InputMediaAudio represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	// Type of the result, must be "audio"
	Type string `json:"type"`
	// File to send as file_id, HTTP URL, or "attach://<file_attach_name>"
	Media string `json:"media"`
	// Optional: Thumbnail of the file in JPEG format
	Thumbnail *InputFileOrString `json:"thumbnail,omitempty"`
	// Optional: Caption of the audio
	Caption string `json:"caption,omitempty"`
	// Optional: Mode for parsing entities in the caption
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional: List of special entities in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional: Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`
	// Optional: Performer of the audio
	Performer string `json:"performer,omitempty"`
	// Optional: Title of the audio
	Title string `json:"title,omitempty"`
}

// InputMediaType returns the type of the media, which is "audio".
func (a *InputMediaAudio) InputMediaType() string {
	return a.Type
}
