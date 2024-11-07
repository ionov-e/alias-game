package types

// BusinessIntro
// Contains information about the start page settings of a Telegram Business account.
type BusinessIntro struct {
	// Optional. Title text of the business intro
	Title *string `json:"title,omitempty"`
	// Optional. Message text of the business intro
	Message *string `json:"message,omitempty"`
	// Optional. Sticker of the business intro
	Sticker *Sticker `json:"sticker,omitempty"`
}
