package types

// ChatPermissions describes the actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`         // Optional. True if the user can send text messages, contacts, giveaways, invoices, locations, and venues.
	CanSendAudios         bool `json:"can_send_audios,omitempty"`           // Optional. True if the user can send audio files.
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`        // Optional. True if the user can send documents.
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`           // Optional. True if the user can send photos.
	CanSendVideos         bool `json:"can_send_videos,omitempty"`           // Optional. True if the user can send videos.
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`      // Optional. True if the user can send video notes.
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`      // Optional. True if the user can send voice notes.
	CanSendPolls          bool `json:"can_send_polls,omitempty"`            // Optional. True if the user can send polls.
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`   // Optional. True if the user can send animations, games, stickers, and use inline bots.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"` // Optional. True if the user can add web page previews to messages.
	CanChangeInfo         bool `json:"can_change_info,omitempty"`           // Optional. True if the user can change the chat title, photo, and other settings (ignored in public supergroups).
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`          // Optional. True if the user can invite new users to the chat.
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`          // Optional. True if the user can pin messages (ignored in public supergroups).
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`         // Optional. True if the user can create forum topics (defaults to CanPinMessages if omitted).
}
