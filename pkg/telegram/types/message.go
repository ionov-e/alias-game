package types

// Message https://core.telegram.org/bots/api#message
type Message struct {
	// unique message identifier inside this chat
	MessageID int64 `json:"message_id"` //nolint:tagliatelle
	// the message belongs to
	Chat Chat `json:"chat"`
	// Optional: sender of the message; may be empty for messages sent to channels.
	User User `json:"from"`
	// Optional: the message was sent in Unix time. It is always a positive number, representing a valid date
	Date uint32 `json:"date"`
	// Optional
	Text          string            `json:"text"`
	Entities      []MessageEntity   `json:"entities"`
	Photo         []PhotoSize       `json:"photo"`
	Audio         Audio             `json:"audio"`
	Video         Video             `json:"video"`
	VideoNote     VideoNote         `json:"video_note"` //nolint:tagliatelle
	Document      Document          `json:"document"`
	Story         Story             `json:"story"`
	Animation     Animation         `json:"animation"`
	Sticker       Sticker           `json:"sticker"`
	ExternalReply ExternalReplyInfo `json:"external_reply"` //nolint:tagliatelle
	Contact       Contact           `json:"contact"`
	Dice          Dice              `json:"dice"`
	Game          Game              `json:"game"`
	Location      Location          `json:"location"`
	Voice         Voice             `json:"voice"`
	Invoice       Invoice           `json:"invoice"`
	Giveaway      Giveaway          `json:"giveaway"`
}
