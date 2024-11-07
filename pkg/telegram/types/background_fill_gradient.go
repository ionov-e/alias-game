package types

// BackgroundFillGradient
// The background is a gradient fill.
type BackgroundFillGradient struct {
	// Type of the background fill, always “gradient”
	Type string `json:"type"`
	// Top color of the gradient in the RGB24 format
	TopColor int `json:"top_color"`
	// Bottom color of the gradient in the RGB24 format
	BottomColor int `json:"bottom_color"`
	// Clockwise rotation angle of the background fill in degrees; 0-359
	RotationAngle int `json:"rotation_angle"`
}

// BackgroundType implements the BackgroundFill interface
func (b BackgroundFillGradient) BackgroundType() string {
	return b.Type
}
