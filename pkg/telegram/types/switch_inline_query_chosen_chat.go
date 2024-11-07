package types

// SwitchInlineQueryChosenChat represents an inline button that switches the user to inline mode in a chosen chat, with an optional default query.
type SwitchInlineQueryChosenChat struct {
	Query             *string `json:"query,omitempty"`               // Optional. Default inline query to insert in the input field.
	AllowUserChats    *bool   `json:"allow_user_chats,omitempty"`    // Optional. True if private chats with users can be chosen.
	AllowBotChats     *bool   `json:"allow_bot_chats,omitempty"`     // Optional. True if private chats with bots can be chosen.
	AllowGroupChats   *bool   `json:"allow_group_chats,omitempty"`   // Optional. True if group and supergroup chats can be chosen.
	AllowChannelChats *bool   `json:"allow_channel_chats,omitempty"` // Optional. True if channel chats can be chosen.
}