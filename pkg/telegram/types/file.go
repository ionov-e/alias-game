package types

// File
// This object represents a file ready to be downloaded.
type File struct {
	// FileID Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID Unique identifier for this file, which is supposed to be the same over time and for different bots
	FileUniqueID string `json:"file_unique_id"`
	// FileSize File size in bytes
	FileSize int64 `json:"file_size"`
	// FilePath File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file
	FilePath string `json:"file_path"`
}
