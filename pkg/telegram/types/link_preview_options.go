package types

// LinkPreviewOptions
// Describes the options used for link preview generation.
type LinkPreviewOptions struct {
	// IsDisabled True, if the link preview is disabled
	IsDisabled bool `json:"is_disabled"`
	// URL to use for the link preview. If empty, the first URL found in the message text will be used
	URL string `json:"url"`
	// PreferSmallMedia True, if the media in the link preview is supposed to be shrunk
	PreferSmallMedia bool `json:"prefer_small_media"`
	// PreferLargeMedia True, if the media in the link preview is supposed to be enlarged
	PreferLargeMedia bool `json:"prefer_large_media"`
	// ShowAboveText True, if the link preview must be shown above the message text
	ShowAboveText bool `json:"show_above_text"`
}
