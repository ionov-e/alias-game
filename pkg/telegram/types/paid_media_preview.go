package types

// PaidMediaPreview
// The paid media isn't available before the payment.
type PaidMediaPreview struct {
	// Type of the paid media, always “preview”
	Type string `json:"type"`
	// Optional. Media width as defined by the sender
	Width *int `json:"width,omitempty"`
	// Optional. Media height as defined by the sender
	Height *int `json:"height,omitempty"`
	// Optional. Duration of the media in seconds as defined by the sender
	Duration *int `json:"duration,omitempty"`
}

// PaidMediaType implements the PaidMedia interface
func (p PaidMediaPreview) PaidMediaType() string {
	return p.Type
}
