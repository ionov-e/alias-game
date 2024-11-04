package types

// Video https://core.telegram.org/bots/api#video
type Video struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Width    uint32 `json:"width"`
	Height   uint32 `json:"height"`
	Duration uint32 `json:"duration"`
}
