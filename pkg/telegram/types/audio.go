package types

// Audio https://core.telegram.org/bots/api#audio represents an audio file, treated as music by Telegram clients.
type Audio struct {
	// Identifier for this file, used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Duration of the audio in seconds.
	Duration int `json:"duration"`
	// Optional. Performer of the audio.
	Performer string `json:"performer,omitempty"`
	// Optional. Title of the audio.
	Title string `json:"title,omitempty"`
	// Optional. Original filename of the audio.
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file.
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
	// Optional. Thumbnail of the album cover.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}
