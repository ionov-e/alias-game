package types

// PhotoSize https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        uint32 `json:"width"`          // Photo width
	Height       uint32 `json:"height"`         // Photo height
	FileSize     *int64 `json:"file_size"`      // Optional. File size in bytes
}
