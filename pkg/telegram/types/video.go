package types

// Video https://core.telegram.org/bots/api#video
type Video struct {
	FileID       string     `json:"file_id"`             // Identifier for this file, used to download or reuse the file.
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier, consistent over time and across bots.
	Width        int        `json:"width"`               // Video width as defined by the sender.
	Height       int        `json:"height"`              // Video height as defined by the sender.
	Duration     int        `json:"duration"`            // Duration of the video in seconds.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Thumbnail of the video.
	FileName     *string    `json:"file_name,omitempty"` // Optional. Original filename of the video.
	MimeType     *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file.
	FileSize     *int64     `json:"file_size,omitempty"` // Optional. File size in bytes.
}
