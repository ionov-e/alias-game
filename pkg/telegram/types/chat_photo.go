package types

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`        // File ID of the small (160x160) chat photo.
	SmallFileUniqueID string `json:"small_file_unique_id"` // Unique file ID of the small chat photo.
	BigFileID         string `json:"big_file_id"`          // File ID of the big (640x640) chat photo.
	BigFileUniqueID   string `json:"big_file_unique_id"`   // Unique file ID of the big chat photo.
}
