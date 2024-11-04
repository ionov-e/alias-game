package types

// VideoNote https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Length   uint32 `json:"length"`
	Duration uint32 `json:"duration"`
}
