package types

// BackgroundTypeFill
// The background is automatically filled based on the selected colors.
type BackgroundTypeFill struct {
	// Type of the background, always "fill"
	Type string `json:"type"`
	// Fill Background fill
	Fill BackgroundFill `json:"fill"`
	// DarkThemeDimming Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`
}

// BackgroundType returns the type of the background, always "fill"
func (b BackgroundTypeFill) BackgroundType() string {
	return b.Type
}
