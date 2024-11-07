package types

// ChatMemberRestricted represents a chat member that is under certain restrictions in the chat.
type ChatMemberRestricted struct {
	Status                string `json:"status"`                    // The member's status in the chat, always “restricted”.
	User                  User   `json:"user"`                      // Information about the user.
	IsMember              bool   `json:"is_member"`                 // True if the user is a member of the chat.
	CanSendMessages       bool   `json:"can_send_messages"`         // True if the user can send text messages and other basic types.
	CanSendAudios         bool   `json:"can_send_audios"`           // True if the user can send audios.
	CanSendDocuments      bool   `json:"can_send_documents"`        // True if the user can send documents.
	CanSendPhotos         bool   `json:"can_send_photos"`           // True if the user can send photos.
	CanSendVideos         bool   `json:"can_send_videos"`           // True if the user can send videos.
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`      // True if the user can send video notes.
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`      // True if the user can send voice notes.
	CanSendPolls          bool   `json:"can_send_polls"`            // True if the user can send polls.
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`   // True if the user can send animations, games, etc.
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"` // True if the user can add web page previews.
	CanChangeInfo         bool   `json:"can_change_info"`           // True if the user can change chat settings.
	CanInviteUsers        bool   `json:"can_invite_users"`          // True if the user can invite new users.
	CanPinMessages        bool   `json:"can_pin_messages"`          // True if the user can pin messages.
	CanManageTopics       bool   `json:"can_manage_topics"`         // True if the user can manage forum topics.
	UntilDate             int64  `json:"until_date"`                // Date when restrictions will be lifted, Unix time.
}

// GetStatus returns the member's status.
func (c ChatMemberRestricted) GetStatus() string {
	return c.Status
}

// GetUser returns information about the user.
func (c ChatMemberRestricted) GetUser() User {
	return c.User
}
