package types

// PaidMediaPhoto
// The paid media is a photo.
type PaidMediaPhoto struct {
	// Type of the paid media, always “photo”
	Type string `json:"type"`
	// The photo
	Photo []PhotoSize `json:"photo"`
}

// PaidMediaType implements the PaidMedia interface
func (p PaidMediaPhoto) PaidMediaType() string {
	return p.Type
}
