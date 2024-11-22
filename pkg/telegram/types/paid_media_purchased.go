package types

type PaidMediaPurchased struct {
	// User who purchased the media
	From User `json:"from"`
	// Bot-specified paid media payload
	PaidMediaPayload string `json:"paid_media_payload"`
}
