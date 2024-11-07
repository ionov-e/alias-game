package types

// InputMediaAudio represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	Type            string           `json:"type"`                       // Type of the result, must be "audio"
	Media           string           `json:"media"`                      // File to send as file_id, HTTP URL, or "attach://<file_attach_name>"
	Thumbnail       *interface{}     `json:"thumbnail,omitempty"`        // Optional: Thumbnail of the file in JPEG format
	Caption         *string          `json:"caption,omitempty"`          // Optional: Caption of the audio
	ParseMode       *string          `json:"parse_mode,omitempty"`       // Optional: Mode for parsing entities in the caption
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"` // Optional: List of special entities in the caption
	Duration        *int             `json:"duration,omitempty"`         // Optional: Duration of the audio in seconds
	Performer       *string          `json:"performer,omitempty"`        // Optional: Performer of the audio
	Title           *string          `json:"title,omitempty"`            // Optional: Title of the audio
}

// InputMediaType returns the type of the media, which is "audio".
func (a *InputMediaAudio) InputMediaType() string {
	return a.Type
}
