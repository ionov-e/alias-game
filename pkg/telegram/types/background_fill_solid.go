package types

// BackgroundFillSolid
// The background is filled using the selected color.
type BackgroundFillSolid struct {
	// Type of the background fill, always “solid”
	Type string `json:"type"`
	// Color of the background fill in the RGB24 format
	Color int `json:"color"`
}

// BackgroundType implements the BackgroundFill interface
func (b BackgroundFillSolid) BackgroundType() string {
	return b.Type
}
