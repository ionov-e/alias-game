package types

// Voice https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID       string  `json:"file_id"`             // Identifier for this file, used to download or reuse the file.
	FileUniqueID string  `json:"file_unique_id"`      // Unique identifier, consistent over time and across bots.
	Duration     int     `json:"duration"`            // Duration of the audio in seconds.
	MimeType     *string `json:"mime_type,omitempty"` // Optional. MIME type of the file.
	FileSize     *int64  `json:"file_size,omitempty"` // Optional. File size in bytes.
}
