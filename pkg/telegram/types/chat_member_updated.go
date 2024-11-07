package types

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`                                  // Chat the user belongs to.
	From                    User            `json:"from"`                                  // Performer of the action that resulted in the change.
	Date                    int64           `json:"date"`                                  // Date of the change in Unix time.
	OldChatMember           ChatMember      `json:"old_chat_member"`                       // Previous information about the chat member.
	NewChatMember           ChatMember      `json:"new_chat_member"`                       // New information about the chat member.
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`                 // Optional. Invite link used by the user to join the chat.
	ViaJoinRequest          *bool           `json:"via_join_request,omitempty"`            // Optional. True if the user joined the chat after a direct join request.
	ViaChatFolderInviteLink *bool           `json:"via_chat_folder_invite_link,omitempty"` // Optional. True if the user joined via a chat folder invite link.
}
