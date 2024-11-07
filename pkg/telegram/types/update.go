package types

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID                uint64                       `json:"update_id"`                           // Unique identifier for the update.
	Message                 *Message                     `json:"message,omitempty"`                   // Optional. New incoming message of any kind - text, photo, sticker, etc.
	EditedMessage           *Message                     `json:"edited_message,omitempty"`            // Optional. New version of a known message that was edited.
	ChannelPost             *Message                     `json:"channel_post,omitempty"`              // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`       // Optional. New version of a known channel post that was edited.
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`       // Optional. The bot was connected to/disconnected from a business account.
	BusinessMessage         *Message                     `json:"business_message,omitempty"`          // Optional. New message from a connected business account.
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`   // Optional. Edited message from a connected business account.
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"` // Optional. Messages deleted from a connected business account.
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`          // Optional. A message reaction was changed.
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`    // Optional. Anonymous reactions to a message were changed.
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`              // Optional. New incoming inline query.
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`      // Optional. Result of an inline query chosen by a user.
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`            // Optional. New incoming callback query.
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`            // Optional. New incoming shipping query.
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`        // Optional. New incoming pre-checkout query.
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`      // Optional. User purchased paid media.
	Poll                    *Poll                        `json:"poll,omitempty"`                      // Optional. New poll state.
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`               // Optional. A user changed their answer in a non-anonymous poll.
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`            // Optional. Bot's chat member status updated.
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`               // Optional. A chat member's status updated.
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`         // Optional. A request to join the chat.
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`                // Optional. A chat boost was added or changed.
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`        // Optional. A boost was removed from the chat.
}
