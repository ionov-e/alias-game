package types

// Document https://core.telegram.org/bots/api#document represents a general file, distinct from specific media types like photos or audio.
type Document struct {
	FileID       string     `json:"file_id"`             // Identifier for this file, used to download or reuse the file.
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier, consistent over time and across bots.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Thumbnail of the document.
	FileName     string     `json:"file_name,omitempty"` // Optional. Original filename of the document.
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file.
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes.
}
