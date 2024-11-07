package types

// VideoChatEnded
// This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	// Duration Video chat duration in seconds
	Duration int64 `json:"duration"`
}
