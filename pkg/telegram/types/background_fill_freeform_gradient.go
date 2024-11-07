package types

// BackgroundFillFreeformGradient
// The background is a freeform gradient that rotates after every message in the chat.
type BackgroundFillFreeformGradient struct {
	// Type of the background fill, always “freeform_gradient”
	Type string `json:"type"`
	// List of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format
	Colors []int `json:"colors"`
}

// BackgroundType implements the BackgroundFill interface
func (b BackgroundFillFreeformGradient) BackgroundType() string {
	return b.Type
}
