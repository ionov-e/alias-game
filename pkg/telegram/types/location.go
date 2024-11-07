package types

// Location represents a point on the map. https://core.telegram.org/bots/api#location
type Location struct {
	Latitude             float64  `json:"latitude"`                         // Latitude as defined by the sender
	Longitude            float64  `json:"longitude"`                        // Longitude as defined by the sender
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           *int     `json:"live_period,omitempty"`            // Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	Heading              *int     `json:"heading,omitempty"`                // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"` // Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}
