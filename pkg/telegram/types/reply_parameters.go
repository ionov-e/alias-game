package types

type ReplyParameters struct {
	MessageID                int             `json:"message_id"`                  // ID of the message being replied to in the current chat
	ChatID                   *string         `json:"chat_id,omitempty"`           // Optional: Chat ID or username if reply is in a different chat
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"` // Optional: True if the reply can be sent without finding the original message
	Quote                    *string         `json:"quote,omitempty"`             // Optional: Quoted part of the message to be replied to
	QuoteParseMode           *string         `json:"quote_parse_mode,omitempty"`  // Optional: Mode for parsing entities in the quote
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`    // Optional: List of special entities in the quote
	QuotePosition            *int            `json:"quote_position,omitempty"`    // Optional: Position of the quote in UTF-16 code units
}
