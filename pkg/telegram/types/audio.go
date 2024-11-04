package types

// Audio https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Duration uint32 `json:"duration"`
}
