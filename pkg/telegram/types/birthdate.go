package types

// Birthdate
// Describes the birthdate of a user.
type Birthdate struct {
	// Day of the user's birth; 1-31
	Day int `json:"day"`
	// Month of the user's birth; 1-12
	Month int `json:"month"`
	// Optional. Year of the user's birth
	Year *int `json:"year,omitempty"`
}
