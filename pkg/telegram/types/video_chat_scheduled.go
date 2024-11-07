package types

// VideoChatScheduled
// This object represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	// StartDate Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
	StartDate int64 `json:"start_date"`
}
