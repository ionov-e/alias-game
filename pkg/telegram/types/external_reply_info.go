package types

// ExternalReplyInfo https://core.telegram.org/bots/api#externalreplyinfo
type ExternalReplyInfo struct {
	// Origin of the message replied to by the given message
	MessageOrigin MessageOrigin `json:"origin"`
	// Optional: chat the original message belongs to. Available only if the chat is a supergroup or a channel
	Chat Chat `json:"chat"`
	// Optional: Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel
	MessageId uint64 `json:"message_id"` //nolint:tagliatelle
}
