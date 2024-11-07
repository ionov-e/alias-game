package types

// BusinessLocation
// Contains information about the location of a Telegram Business account.
type BusinessLocation struct {
	// Address of the business
	Address string `json:"address"`
	// Optional. Location of the business
	Location *Location `json:"location,omitempty"`
}
