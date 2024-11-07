package types

// VideoNote https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID       string     `json:"file_id"`             // Identifier for this file, used to download or reuse the file.
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier, consistent over time and across bots.
	Length       int        `json:"length"`              // Diameter of the video message.
	Duration     int        `json:"duration"`            // Duration of the video in seconds.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Thumbnail of the video note.
	FileSize     *int64     `json:"file_size,omitempty"` // Optional. File size in bytes.
}
