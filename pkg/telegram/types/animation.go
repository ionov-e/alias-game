package types

// Animation https://core.telegram.org/bots/api#animation represents an animation file, such as a GIF or a video without sound.
type Animation struct {
	// Identifier for this file, used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Video width as defined by the sender.
	Width int `json:"width"`
	// Video height as defined by the sender.
	Height int `json:"height"`
	// Duration of the video in seconds.
	Duration int `json:"duration"`
	// Optional. Thumbnail of the animation.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// Optional. Original filename of the animation.
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by the sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
}
