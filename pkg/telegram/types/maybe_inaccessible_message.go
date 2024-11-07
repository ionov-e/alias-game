package types

// MaybeInaccessibleMessage represents a message that can either be a regular or an inaccessible message.
type MaybeInaccessibleMessage interface {
	MessageMessageID() int64
	MessageChat() Chat
	MessageDate() int64
}
