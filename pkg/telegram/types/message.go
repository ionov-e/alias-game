package types

// Message https://core.telegram.org/bots/api#message
type Message struct {
	// Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageID int64 `json:"message_id"`
	// Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// Optional. Sender of the message; may be empty for messages sent to channels. For backward compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in non-channel chats
	From *User `json:"from,omitempty"`
	// Optional. Sender of the message when sent on behalf of a chat. For example, the supergroup itself for messages sent by its anonymous administrators or a linked channel for messages automatically forwarded to the channel's discussion group. For backward compatibility, if the message was sent on behalf of a chat, the field from contains a fake sender user in non-channel chats.
	SenderChat *Chat `json:"sender_chat,omitempty"`
	// Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBoostCount int `json:"sender_boost_count,omitempty"`
	// Optional. The bot that actually sent the message on behalf of the business account. Available only for outgoing messages sent on behalf of the connected business account.
	SenderBusinessBot *User `json:"sender_business_bot,omitempty"`
	// Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	Date int64 `json:"date"`
	// Optional. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent from any potential bot chat which might share the same identifier.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Chat the message belongs to.
	Chat Chat `json:"chat"`
	// Optional. Information about the original message for forwarded messages
	ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`
	// Optional. True, if the message is sent to a forum topic
	IsTopicMessage bool `json:"is_topic_message,omitempty"`
	// Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`
	// Optional. For replies in the same chat and message thread, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	ExternalReply *ExternalReplyInfo `json:"external_reply,omitempty"`
	// Optional. For replies that quote part of the original message, the quoted part of the message
	Quote *TextQuote `json:"quote,omitempty"`
	// Optional. For replies to a story, the original story
	ReplyToStory *Story `json:"reply_to_story,omitempty"`
	// Optional. Bot that sent the message.
	ViaBot *User `json:"via_bot,omitempty"`
	// Optional. Date the message was last edited in Unix time
	EditDate int64 `json:"edit_date,omitempty"`
	// Optional. True if message cannot be forwarded.
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// Optional. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message
	IsFromOffline bool `json:"is_from_offline,omitempty"`
	// Optional. ID for media message group.
	MediaGroupID string `json:"media_group_id,omitempty"`
	// Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`
	// Optional. Actual text of the message.
	Text string `json:"text,omitempty"`
	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`
	// Optional. Options used for link preview generation for the message, if it is a text message and link preview options were changed
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	// Optional. Unique ID for message effect.
	EffectID string `json:"effect_id,omitempty"`
	// Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`
	// Optional. Information about audio file.
	Audio *Audio `json:"audio,omitempty"`
	// Optional. Information about general file.
	Document *Document `json:"document,omitempty"`
	// Optional. Information about paid media.
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`
	// Optional. Available sizes of photo.
	Photo []PhotoSize `json:"photo,omitempty"`
	// Optional. Information about sticker.
	Sticker *Sticker `json:"sticker,omitempty"`
	// Optional. Forwarded story.
	Story *Story `json:"story,omitempty"`
	// Optional. Information about video.
	Video *Video `json:"video,omitempty"`
	// Optional. Information about video message.
	VideoNote *VideoNote `json:"video_note,omitempty"`
	// Optional. Information about voice message.
	Voice *Voice `json:"voice,omitempty"`
	// Optional. Caption for media.
	Caption string `json:"caption,omitempty"`
	// Optional. Special entities in the caption.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional. True if caption must be shown above media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`
	// Optional. True if media is covered by a spoiler.
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`
	// Optional. Information about shared contact.
	Contact *Contact `json:"contact,omitempty"`
	// Optional. Dice with random value.
	Dice *Dice `json:"dice,omitempty"`
	// Optional. Information about game.
	Game *Game `json:"game,omitempty"`
	// Optional. Information about poll.
	Poll *Poll `json:"poll,omitempty"`
	// Optional. Information about venue.
	Venue *Venue `json:"venue,omitempty"`
	// Optional. Information about location.
	Location *Location `json:"location,omitempty"`
	// Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	NewChatMembers []User `json:"new_chat_members,omitempty"`
	// Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// Optional. New chat title.
	NewChatTitle string `json:"new_chat_title,omitempty"`
	// Optional. New chat photo.
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`
	// Optional. Chat photo deleted.
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// Optional. Group created.
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
	// Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	// Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
	// Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	// Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	// Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`
	// Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Optional. Invoice for payment.
	Invoice *Invoice `json:"invoice,omitempty"`
	// Optional. Information about successful payment.
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	// Optional. Information about refunded payment.
	RefundedPayment *RefundedPayment `json:"refunded_payment,omitempty"`
	// Optional. Service message for shared users.
	UsersShared *UsersShared `json:"users_shared,omitempty"`
	// Optional. Service message for shared chat.
	ChatShared *ChatShared `json:"chat_shared,omitempty"`
	// Optional. Domain of connected website.
	ConnectedWebsite string `json:"connected_website,omitempty"`
	// Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess: https://core.telegram.org/bots/webapps#initializing-mini-apps
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`
	// Optional. Telegram Passport data.
	PassportData *PassportData `json:"passport_data,omitempty"`
	// Optional. User triggered another user's proximity alert.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	// Optional. Boost added service message.
	BoostAdded *ChatBoostAdded `json:"boost_added,omitempty"`
	// Optional. Chat background set.
	ChatBackgroundSet *ChatBackground `json:"chat_background_set,omitempty"`
	// Optional. Forum topic created.
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`
	// Optional. Forum topic edited.
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`
	// Optional. Forum topic closed.
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed,omitempty"`
	// Optional. Forum topic reopened.
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened,omitempty"`
	// Optional. General forum topic hidden.
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`
	// Optional. General forum topic unhidden.
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`
	// Optional. Giveaway created.
	GiveawayCreated *GiveawayCreated `json:"giveaway_created,omitempty"`
	// Optional. Scheduled giveaway message.
	Giveaway *Giveaway `json:"giveaway,omitempty"`
	// Optional. Giveaway with public winners completed.
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`
	// Optional. Giveaway without public winners completed.
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed,omitempty"`
	// Optional. Video chat scheduled.
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
	// Optional. Video chat started.
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`
	// Optional. Video chat ended.
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`
	// Optional. New video chat participants invited.
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	// Optional. Data sent by a Web App.
	WebAppData *WebAppData `json:"web_app_data,omitempty"`
	// Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
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
