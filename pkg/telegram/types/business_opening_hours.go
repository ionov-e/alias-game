package types

// BusinessOpeningHours
// Describes the opening hours of a business.
type BusinessOpeningHours struct {
	// Unique name of the time zone for which the opening hours are defined
	TimeZoneName string `json:"time_zone_name"`
	// List of time intervals describing business opening hours
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}
