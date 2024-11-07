package types

// ChatFullInfo represents full information about a chat.
type ChatFullInfo struct {
	Chat
	AccentColorID                      int                   `json:"accent_color_id"`                                   // Accent color ID for chat name and backgrounds.
	MaxReactionCount                   int                   `json:"max_reaction_count"`                                // Maximum reactions allowed in the chat.
	Photo                              *ChatPhoto            `json:"photo,omitempty"`                                   // Optional. Chat photo.
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`                        // Optional. List of active usernames in the chat.
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`                               // Optional. Birthdate for private chats.
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`                          // Optional. Intro of the business in private chats with business accounts.
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`                       // Optional. Business location for private chats.
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`                  // Optional. Business opening hours.
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`                           // Optional. Personal channel for private chats.
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`                     // Optional. List of allowed reactions in the chat.
	BackgroundCustomEmojiID            *string               `json:"background_custom_emoji_id,omitempty"`              // Optional. Emoji ID for reply header and link preview background.
	ProfileAccentColorID               *int                  `json:"profile_accent_color_id,omitempty"`                 // Optional. Accent color ID for profile background.
	ProfileBackgroundCustomEmojiID     *string               `json:"profile_background_custom_emoji_id,omitempty"`      // Optional. Emoji ID for profile background.
	EmojiStatusCustomEmojiID           *string               `json:"emoji_status_custom_emoji_id,omitempty"`            // Optional. Emoji status ID in private chat.
	EmojiStatusExpirationDate          *int64                `json:"emoji_status_expiration_date,omitempty"`            // Optional. Expiration date for emoji status in Unix time.
	Bio                                *string               `json:"bio,omitempty"`                                     // Optional. Bio of the other party in private chat.
	HasPrivateForwards                 *bool                 `json:"has_private_forwards,omitempty"`                    // Optional. True if links are restricted to chats with the user.
	HasRestrictedVoiceAndVideoMessages *bool                 `json:"has_restricted_voice_and_video_messages,omitempty"` // Optional. True if voice/video note messages are restricted.
	JoinToSendMessages                 *bool                 `json:"join_to_send_messages,omitempty"`                   // Optional. True if users must join the supergroup to send messages.
	JoinByRequest                      *bool                 `json:"join_by_request,omitempty"`                         // Optional. True if join requests need approval.
	Description                        *string               `json:"description,omitempty"`                             // Optional. Description for groups, supergroups, and channels.
	InviteLink                         *string               `json:"invite_link,omitempty"`                             // Optional. Primary invite link.
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`                          // Optional. Most recent pinned message.
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`                             // Optional. Default permissions for group/supergroup members.
	CanSendPaidMedia                   *bool                 `json:"can_send_paid_media,omitempty"`                     // Optional. True if paid media messages are allowed.
	SlowModeDelay                      *int                  `json:"slow_mode_delay,omitempty"`                         // Optional. Minimum delay between consecutive messages.
	UnrestrictBoostCount               *int                  `json:"unrestrict_boost_count,omitempty"`                  // Optional. For supergroups, the minimum number of boosts that a non-administrator user needs to add in order to ignore slow mode and chat permissions
	MessageAutoDeleteTime              *int                  `json:"message_auto_delete_time,omitempty"`                // Optional. Message auto-delete time in seconds.
	HasAggressiveAntiSpamEnabled       *bool                 `json:"has_aggressive_anti_spam_enabled,omitempty"`        // Optional. True if aggressive anti-spam checks are enabled.
	HasHiddenMembers                   *bool                 `json:"has_hidden_members,omitempty"`                      // Optional. True if only administrators can view all members.
	HasProtectedContent                *bool                 `json:"has_protected_content,omitempty"`                   // Optional. True if messages can't be forwarded.
	HasVisibleHistory                  *bool                 `json:"has_visible_history,omitempty"`                     // Optional. True if new members can access old messages.
	StickerSetName                     *string               `json:"sticker_set_name,omitempty"`                        // Optional. Name of group sticker set for supergroups.
	CanSetStickerSet                   *bool                 `json:"can_set_sticker_set,omitempty"`                     // Optional. True if the bot can change the sticker set.
	CustomEmojiStickerSetName          *string               `json:"custom_emoji_sticker_set_name,omitempty"`           // Optional. Name of group's custom emoji sticker set.
	LinkedChatID                       *int64                `json:"linked_chat_id,omitempty"`                          // Optional. Unique ID for linked chat.
	Location                           *ChatLocation         `json:"location,omitempty"`                                // Optional. Location of the supergroup.
}
