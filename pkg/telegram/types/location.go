package types

// Location https://core.telegram.org/bots/api#location
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
