package types

// KeyboardButtonRequestChat defines criteria to request a suitable chat.
type KeyboardButtonRequestChat struct {
	// Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
	RequestID int32 `json:"request_id"`
	// Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsChannel bool `json:"chat_is_channel"`
	// Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
	ChatIsForum bool `json:"chat_is_forum,omitempty"`
	// Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
	ChatHasUsername bool `json:"chat_has_username,omitempty"`
	// Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	ChatIsCreated bool `json:"chat_is_created,omitempty"`
	// Optional. A JSON-serialized object listing the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	// Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
	BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	// Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	BotIsMember bool `json:"bot_is_member,omitempty"`
	// Optional. Pass True to request the chat's title.
	RequestTitle bool `json:"request_title,omitempty"`
	// Optional. Pass True to request the chat's username.
	RequestUsername bool `json:"request_username,omitempty"`
	// Optional. Pass True to request the chat's photo.
	RequestPhoto bool `json:"request_photo,omitempty"`
}
