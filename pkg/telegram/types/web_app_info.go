package types

// WebAppInfo describes a Web App.
type WebAppInfo struct {
	// An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
	URL string `json:"url"`
}
