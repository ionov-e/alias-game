package types

// PaidMediaInfo
// Describes the paid media added to a message.
type PaidMediaInfo struct {
	// The number of Telegram Stars that must be paid to buy access to the media
	StarCount int `json:"star_count"`
	// Information about the paid media
	PaidMedia []PaidMedia `json:"paid_media"`
}
