package types

// VideoNote https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	// Identifier for this file, used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier, consistent over time and across bots.
	FileUniqueID string `json:"file_unique_id"`
	// Video width and height (diameter of the video message) as defined by the sender
	Length int `json:"length"`
	// Duration of the video in seconds.
	Duration int `json:"duration"`
	// Optional. Thumbnail of the video note.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// Optional. File size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}
