package types

// ChatMemberAdministrator represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	Status              string `json:"status"`                      // The member's status in the chat, always “administrator”.
	User                User   `json:"user"`                        // Information about the user.
	CanBeEdited         bool   `json:"can_be_edited"`               // True if the bot can edit admin privileges.
	IsAnonymous         bool   `json:"is_anonymous"`                // True if the user's presence is hidden.
	CanManageChat       bool   `json:"can_manage_chat"`             // True if the admin can manage chat.
	CanDeleteMessages   bool   `json:"can_delete_messages"`         // True if the admin can delete messages.
	CanManageVideoChats bool   `json:"can_manage_video_chats"`      // True if the admin can manage video chats.
	CanRestrictMembers  bool   `json:"can_restrict_members"`        // True if the admin can restrict or ban members.
	CanPromoteMembers   bool   `json:"can_promote_members"`         // True if the admin can promote members.
	CanChangeInfo       bool   `json:"can_change_info"`             // True if the admin can change chat settings.
	CanInviteUsers      bool   `json:"can_invite_users"`            // True if the admin can invite users.
	CanPostStories      bool   `json:"can_post_stories"`            // True if the admin can post stories.
	CanEditStories      bool   `json:"can_edit_stories"`            // True if the admin can edit stories.
	CanDeleteStories    bool   `json:"can_delete_stories"`          // True if the admin can delete stories.
	CanPostMessages     *bool  `json:"can_post_messages,omitempty"` // Optional. True if the admin can post messages in the channel.
	CanEditMessages     *bool  `json:"can_edit_messages,omitempty"` // Optional. True if the admin can edit messages.
	CanPinMessages      *bool  `json:"can_pin_messages,omitempty"`  // Optional. True if the admin can pin messages.
	CanManageTopics     *bool  `json:"can_manage_topics,omitempty"` // Optional. True if the admin can manage forum topics.
	CustomTitle         string `json:"custom_title,omitempty"`      // Optional. Custom title for this user.
}

// GetStatus returns the member's status.
func (c ChatMemberAdministrator) GetStatus() string {
	return c.Status
}

// GetUser returns information about the user.
func (c ChatMemberAdministrator) GetUser() User {
	return c.User
}
