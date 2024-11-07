package types

import "fmt"

// ChatMember represents a member of a chat. This is the base type for other specific chat member types. https://core.telegram.org/bots/api#chatmember
type ChatMember interface {
	GetStatus() string
	GetUser() User
}

func ChatMemberFactory(status string, data map[string]interface{}) (ChatMember, error) {
	switch status {
	case "creator":
		return ChatMemberOwner{
			Status:      "creator",
			User:        data["user"].(User),
			IsAnonymous: data["is_anonymous"].(bool),
			CustomTitle: data["custom_title"].(string),
		}, nil
	case "administrator":
		return ChatMemberAdministrator{
			Status:              "administrator",
			User:                data["user"].(User),
			CanBeEdited:         data["can_be_edited"].(bool),
			IsAnonymous:         data["is_anonymous"].(bool),
			CanManageChat:       data["can_manage_chat"].(bool),
			CanDeleteMessages:   data["can_delete_messages"].(bool),
			CanManageVideoChats: data["can_manage_video_chats"].(bool),
			CanRestrictMembers:  data["can_restrict_members"].(bool),
			CanPromoteMembers:   data["can_promote_members"].(bool),
			CanChangeInfo:       data["can_change_info"].(bool),
			CanInviteUsers:      data["can_invite_users"].(bool),
			CanPostStories:      data["can_post_stories"].(bool),
			CanEditStories:      data["can_edit_stories"].(bool),
			CanDeleteStories:    data["can_delete_stories"].(bool),
			CanPostMessages:     boolPtr(data["can_post_messages"].(bool)),
			CanEditMessages:     boolPtr(data["can_edit_messages"].(bool)),
			CanPinMessages:      boolPtr(data["can_pin_messages"].(bool)),
			CanManageTopics:     boolPtr(data["can_manage_topics"].(bool)),
			CustomTitle:         data["custom_title"].(string),
		}, nil
	case "member":
		return ChatMemberMember{
			Status:    "member",
			User:      data["user"].(User),
			UntilDate: int64Ptr(int64(data["until_date"].(float64))),
		}, nil
	case "restricted":
		return ChatMemberRestricted{
			Status:                "restricted",
			User:                  data["user"].(User),
			IsMember:              data["is_member"].(bool),
			CanSendMessages:       data["can_send_messages"].(bool),
			CanSendAudios:         data["can_send_audios"].(bool),
			CanSendDocuments:      data["can_send_documents"].(bool),
			CanSendPhotos:         data["can_send_photos"].(bool),
			CanSendVideos:         data["can_send_videos"].(bool),
			CanSendVideoNotes:     data["can_send_video_notes"].(bool),
			CanSendVoiceNotes:     data["can_send_voice_notes"].(bool),
			CanSendPolls:          data["can_send_polls"].(bool),
			CanSendOtherMessages:  data["can_send_other_messages"].(bool),
			CanAddWebPagePreviews: data["can_add_web_page_previews"].(bool),
			CanChangeInfo:         data["can_change_info"].(bool),
			CanInviteUsers:        data["can_invite_users"].(bool),
			CanPinMessages:        data["can_pin_messages"].(bool),
			CanManageTopics:       data["can_manage_topics"].(bool),
			UntilDate:             int64(data["until_date"].(float64)),
		}, nil
	case "left":
		return ChatMemberLeft{
			Status: "left",
			User:   data["user"].(User),
		}, nil
	case "kicked":
		return ChatMemberBanned{
			Status:    "kicked",
			User:      data["user"].(User),
			UntilDate: int64(data["until_date"].(float64)),
		}, nil
	default:
		return nil, fmt.Errorf("unknown ChatMember status: %s", status)
	}
}

func int64Ptr(i int64) *int64 {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}
