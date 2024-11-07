package types

// BusinessOpeningHoursInterval
// Describes an interval of time during which a business is open.
type BusinessOpeningHoursInterval struct {
	// The minute's sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 * 24 * 60
	OpeningMinute int `json:"opening_minute"`
	// The minute's sequence number in a week, starting on Monday, marking the end of the time interval during which the business is open; 0 - 8 * 24 * 60
	ClosingMinute int `json:"closing_minute"`
}
