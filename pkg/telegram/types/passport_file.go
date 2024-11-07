package types

// PassportFile
// This object represents a file uploaded to Telegram Passport.
type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	FileUniqueID string `json:"file_unique_id"`
	// File size in bytes
	FileSize int64 `json:"file_size"`
	// Unix time when the file was uploaded
	FileDate int64 `json:"file_date"`
}
