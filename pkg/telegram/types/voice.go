package types

// Voice https://core.telegram.org/bots/api#voice
type Voice struct {
	// Identifier for this file, used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Duration of the audio in seconds.
	Duration int `json:"duration"`
	// Optional. MIME type of the file.
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}
