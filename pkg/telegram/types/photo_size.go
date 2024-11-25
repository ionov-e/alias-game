package types

// PhotoSize https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Photo width
	Width uint32 `json:"width"`
	// Photo height
	Height uint32 `json:"height"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size"`
}
