package types

// BackgroundTypeWallpaper
// The background is a wallpaper in the JPEG format.
type BackgroundTypeWallpaper struct {
	// Type of the background, always "wallpaper"
	Type string `json:"type"`
	// Document with the wallpaper
	Document Document `json:"document"`
	// DarkThemeDimming Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`
	// Optional. IsBlurred True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	IsBlurred *bool `json:"is_blurred"`
	// Optional. IsMoving True, if the background moves slightly when the device is tilted
	IsMoving *bool `json:"is_moving"`
}

// BackgroundType returns the type of the background, always "wallpaper"
func (b BackgroundTypeWallpaper) BackgroundType() string {
	return b.Type
}
