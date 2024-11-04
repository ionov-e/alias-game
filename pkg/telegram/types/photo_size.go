package types

// PhotoSize https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Width    uint32 `json:"width"`
	Height   uint32 `json:"height"`
	FileSize uint32 `json:"file_size"` //nolint:tagliatelle
}
