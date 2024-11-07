package types

// InaccessibleMessage represents a message that is inaccessible to the bot.
type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`       // Chat the message belonged to
	MessageID int64 `json:"message_id"` // Unique message identifier inside the chat
	Date      int64 `json:"date"`       // Always 0 to differentiate it from accessible messages
}

// MessageMessageID returns the message ID.
func (im InaccessibleMessage) MessageMessageID() int64 {
	return im.MessageID
}

// MessageChat returns the chat the message belonged to.
func (im InaccessibleMessage) MessageChat() Chat {
	return im.Chat
}

func (im InaccessibleMessage) MessageDate() int64 {
	return im.Date
}
