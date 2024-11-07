package types

// Message https://core.telegram.org/bots/api#message
type Message struct {
	MessageID                     int64                          `json:"message_id"`                                  // Unique ID for this message.
	MessageThreadID               int64                          `json:"message_thread_id,omitempty"`                 // Optional. Thread ID the message belongs to, for supergroups.
	From                          *User                          `json:"from,omitempty"`                              // Optional. Sender; empty if sent to channels.
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`                       // Optional. Sender of message when sent on behalf of a chat.
	SenderBoostCount              int                            `json:"sender_boost_count,omitempty"`                // Optional. Boost count if sender boosted the chat.
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`               // Optional. Bot sending message for business account.
	Date                          int64                          `json:"date"`                                        // Date the message was sent in Unix time.
	BusinessConnectionID          string                         `json:"business_connection_id,omitempty"`            // Optional. ID of business connection.
	Chat                          Chat                           `json:"chat"`                                        // Chat the message belongs to.
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`                    // Optional. Information for forwarded messages.
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`                  // Optional. True if sent to a forum topic.
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`              // Optional. True if channel post auto-forwarded.
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`                  // Optional. Original message in reply.
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`                    // Optional. Information on replied message from another chat.
	Quote                         *TextQuote                     `json:"quote,omitempty"`                             // Optional. Quoted text for replies.
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`                    // Optional. Original story in reply.
	ViaBot                        *User                          `json:"via_bot,omitempty"`                           // Optional. Bot that sent the message.
	EditDate                      int64                          `json:"edit_date,omitempty"`                         // Optional. Last edit date in Unix time.
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`             // Optional. True if message cannot be forwarded.
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`                   // Optional. True if sent as an implicit action.
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`                    // Optional. ID for media message group.
	AuthorSignature               string                         `json:"author_signature,omitempty"`                  // Optional. Signature for posts in channels.
	Text                          string                         `json:"text,omitempty"`                              // Optional. Actual text of the message.
	Entities                      []MessageEntity                `json:"entities,omitempty"`                          // Optional. Special entities in the text.
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`              // Optional. Link preview options.
	EffectID                      string                         `json:"effect_id,omitempty"`                         // Optional. Unique ID for message effect.
	Animation                     *Animation                     `json:"animation,omitempty"`                         // Optional. Information about animation.
	Audio                         *Audio                         `json:"audio,omitempty"`                             // Optional. Information about audio file.
	Document                      *Document                      `json:"document,omitempty"`                          // Optional. Information about general file.
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`                        // Optional. Information about paid media.
	Photo                         []PhotoSize                    `json:"photo,omitempty"`                             // Optional. Available sizes of photo.
	Sticker                       *Sticker                       `json:"sticker,omitempty"`                           // Optional. Information about sticker.
	Story                         *Story                         `json:"story,omitempty"`                             // Optional. Forwarded story.
	Video                         *Video                         `json:"video,omitempty"`                             // Optional. Information about video.
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`                        // Optional. Information about video message.
	Voice                         *Voice                         `json:"voice,omitempty"`                             // Optional. Information about voice message.
	Caption                       string                         `json:"caption,omitempty"`                           // Optional. Caption for media.
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`                  // Optional. Special entities in the caption.
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`          // Optional. True if caption must be shown above media.
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`                 // Optional. True if media is covered by a spoiler.
	Contact                       *Contact                       `json:"contact,omitempty"`                           // Optional. Information about shared contact.
	Dice                          *Dice                          `json:"dice,omitempty"`                              // Optional. Dice with random value.
	Game                          *Game                          `json:"game,omitempty"`                              // Optional. Information about game.
	Poll                          *Poll                          `json:"poll,omitempty"`                              // Optional. Information about poll.
	Venue                         *Venue                         `json:"venue,omitempty"`                             // Optional. Information about venue.
	Location                      *Location                      `json:"location,omitempty"`                          // Optional. Information about location.
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`                  // Optional. New group members.
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`                  // Optional. Member removed from group.
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`                    // Optional. New chat title.
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`                    // Optional. New chat photo.
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`                 // Optional. Chat photo deleted.
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`                // Optional. Group created.
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`           // Optional. Supergroup created.
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`              // Optional. Channel created.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"` // Optional. Auto-delete timer changed.
	MigrateToChatID               int64                          `json:"migrate_to_chat_id,omitempty"`                // Optional. Migrated to supergroup ID.
	MigrateFromChatID             int64                          `json:"migrate_from_chat_id,omitempty"`              // Optional. Migrated from group ID.
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`                    // Optional. Pinned message.
	Invoice                       *Invoice                       `json:"invoice,omitempty"`                           // Optional. Invoice for payment.
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`                // Optional. Information about successful payment.
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`                  // Optional. Information about refunded payment.
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`                      // Optional. Service message for shared users.
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`                       // Optional. Service message for shared chat.
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`                 // Optional. Domain of connected website.
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`              // Optional. Write access allowed.
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`                     // Optional. Telegram Passport data.
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`         // Optional. User triggered another user's proximity alert.
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`                       // Optional. Boost added service message.
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`               // Optional. Chat background set.
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`               // Optional. Forum topic created.
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`                // Optional. Forum topic edited.
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`                // Optional. Forum topic closed.
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`              // Optional. Forum topic reopened.
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`        // Optional. General forum topic hidden.
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`      // Optional. General forum topic unhidden.
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`                  // Optional. Giveaway created.
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`                          // Optional. Scheduled giveaway message.
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`                  // Optional. Giveaway with public winners completed.
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`                // Optional. Giveaway without public winners completed.
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`              // Optional. Video chat scheduled.
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`                // Optional. Video chat started.
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`                  // Optional. Video chat ended.
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`   // Optional. New video chat participants invited.
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`                      // Optional. Data sent by a Web App.
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`                      // Optional. Inline keyboard attached to the message.
}

// MessageMessageID returns the message ID.
func (m Message) MessageMessageID() int64 {
	return m.MessageID
}

// MessageChat returns the chat the message belonged to.
func (m Message) MessageChat() Chat {
	return m.Chat
}

func (m Message) MessageDate() int64 {
	return m.Date
}
