package types

// UserProfilePhotos
// This object represents a user's profile pictures.
type UserProfilePhotos struct {
	// TotalCount Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`
	// Photos Requested profile pictures (in up to 4 sizes each)
	Photos [][]PhotoSize `json:"photos"`
}
