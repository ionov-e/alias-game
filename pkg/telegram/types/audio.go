package types

// Audio https://core.telegram.org/bots/api#audio represents an audio file, treated as music by Telegram clients.
type Audio struct {
	FileID       string     `json:"file_id"`             // Identifier for this file, used to download or reuse the file.
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier, consistent over time and across bots.
	Duration     int        `json:"duration"`            // Duration of the audio in seconds.
	Performer    *string    `json:"performer,omitempty"` // Optional. Performer of the audio.
	Title        *string    `json:"title,omitempty"`     // Optional. Title of the audio.
	FileName     *string    `json:"file_name,omitempty"` // Optional. Original filename of the audio.
	MimeType     *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file.
	FileSize     *int64     `json:"file_size,omitempty"` // Optional. File size in bytes.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Thumbnail of the album cover.
}
