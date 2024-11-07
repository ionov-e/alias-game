package types

type PaidMediaPurchased struct {
	From             User   `json:"from"`               // User who purchased the media
	PaidMediaPayload string `json:"paid_media_payload"` // Bot-specified paid media payload
}
