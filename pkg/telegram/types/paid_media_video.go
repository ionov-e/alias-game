package types

// PaidMediaVideo
// The paid media is a video.
type PaidMediaVideo struct {
	// Type of the paid media, always “video”
	Type string `json:"type"`
	// The video
	Video Video `json:"video"`
}

// PaidMediaType implements the PaidMedia interface
func (p PaidMediaVideo) PaidMediaType() string {
	return p.Type
}
