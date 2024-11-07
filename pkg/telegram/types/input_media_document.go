package types

// InputMediaDocument represents a general file to be sent.
type InputMediaDocument struct {
	Type                        string           `json:"type"`                                     // Type of the result, must be "document"
	Media                       string           `json:"media"`                                    // File to send as file_id, HTTP URL, or "attach://<file_attach_name>"
	Thumbnail                   *interface{}     `json:"thumbnail,omitempty"`                      // Optional: Thumbnail of the file in JPEG format
	Caption                     *string          `json:"caption,omitempty"`                        // Optional: Caption of the document
	ParseMode                   *string          `json:"parse_mode,omitempty"`                     // Optional: Mode for parsing entities in the caption
	CaptionEntities             *[]MessageEntity `json:"caption_entities,omitempty"`               // Optional: List of special entities in the caption
	DisableContentTypeDetection *bool            `json:"disable_content_type_detection,omitempty"` // Optional: Disable content type detection
}

// InputMediaType returns the type of the media, which is "document".
func (d *InputMediaDocument) InputMediaType() string {
	return d.Type
}
