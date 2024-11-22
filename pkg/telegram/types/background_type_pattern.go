package types

// BackgroundTypePattern
// The background is a PNG or TGV (gzipped subset of SVG with MIME type “application/x-tgwallpattern”) pattern to be combined with the background fill chosen by the user.
type BackgroundTypePattern struct {
	// Type of the background, always "pattern"
	Type string `json:"type"`
	// Document with the pattern
	Document Document `json:"document"`
	// Fill Background fill that is combined with the pattern
	Fill BackgroundFill `json:"fill"`
	// Intensity of the pattern when it is shown above the filled background; 0-100
	Intensity int `json:"intensity"`
	// Optional. IsInverted True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	IsInverted bool `json:"is_inverted"`
	// Optional. IsMoving True, if the background moves slightly when the device is tilted
	IsMoving bool `json:"is_moving"`
}

// BackgroundType returns the type of the background, always "pattern"
func (b BackgroundTypePattern) BackgroundType() string {
	return b.Type
}
