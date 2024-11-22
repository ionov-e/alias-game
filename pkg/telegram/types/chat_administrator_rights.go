package types

// ChatAdministratorRights represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`
	// True, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`
	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`
	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	// True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanRestrictMembers bool `json:"can_restrict_members"`
	// True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`
	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`
	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`
	// True, if the administrator can post stories to the chat
	CanPostStories bool `json:"can_post_stories"`
	// True, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories bool `json:"can_edit_stories"`
	// True, if the administrator can delete stories posted by other users
	CanDeleteStories bool `json:"can_delete_stories"`
	// Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	// Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CanManageTopics bool `json:"can_manage_topics,omitempty"`
}
