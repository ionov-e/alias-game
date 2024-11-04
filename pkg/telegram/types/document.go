package types

// Document https://core.telegram.org/bots/api#document
type Document struct {
	FileID string `json:"file_id"` //nolint:tagliatelle
}
