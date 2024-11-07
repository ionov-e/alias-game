package types

// VideoChatParticipantsInvited
// This object represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	// Users New members that were invited to the video chat
	Users []User `json:"users"`
}
