package types

// Voice https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Duration uint32 `json:"duration"`
}
