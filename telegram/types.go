package telegram

// MessageResponse https://core.telegram.org/bots/api#sendmessage
type MessageResponse struct {
	Ok          bool    `json:"ok"`
	Description string  `json:"description"`
	Result      Message `json:"result"`
}

func (m *MessageResponse) IsOk() bool {
	return m.Ok
}

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID uint64 `json:"update_id"` //nolint:tagliatelle
	// Optional: new incoming message of any kind - text, photo, sticker, etc.
	Message Message `json:"message"`
	// Optional: new message from a connected business account
	BusinessMessage Message `json:"business_message"` //nolint:tagliatelle
}

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

// Chat https://core.telegram.org/bots/api#chat
type Chat struct {
	// Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it
	ID int64 `json:"id"`
	// Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// Optional: for private chats, supergroups and channels if available
	Username string `json:"username"`
}

// User https://core.telegram.org/bots/api#user
type User struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`     //nolint:tagliatelle
	FirstName string `json:"first_name"` //nolint:tagliatelle
	// Optional: IETF language tag
	Language string `json:"language_code"` //nolint:tagliatelle
}

func (u User) LanguageWithDefault() string {
	if u.Language != "" {
		return u.Language
	}
	return "ru" // TODO check if correct string
}

// MessageEntity https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type string `json:"type"`
	User User   `json:"user"`
}

// ExternalReplyInfo https://core.telegram.org/bots/api#externalreplyinfo
type ExternalReplyInfo struct {
	// Origin of the message replied to by the given message
	MessageOrigin MessageOrigin `json:"origin"`
	// Optional: chat the original message belongs to. Available only if the chat is a supergroup or a channel
	Chat Chat `json:"chat"`
	// Optional: Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel
	MessageId uint64 `json:"message_id"` //nolint:tagliatelle
}

// MessageOrigin https://core.telegram.org/bots/api#messageorigin
type MessageOrigin struct {
	Type string `json:"type"`
	Date uint32 `json:"date"`
}

// PhotoSize https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Width    uint32 `json:"width"`
	Height   uint32 `json:"height"`
	FileSize uint32 `json:"file_size"` //nolint:tagliatelle
}

// Audio https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Duration uint32 `json:"duration"`
}

// Document https://core.telegram.org/bots/api#document
type Document struct {
	FileID string `json:"file_id"` //nolint:tagliatelle
}

// Video https://core.telegram.org/bots/api#video
type Video struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Width    uint32 `json:"width"`
	Height   uint32 `json:"height"`
	Duration uint32 `json:"duration"`
}

// VideoNote https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Length   uint32 `json:"length"`
	Duration uint32 `json:"duration"`
}

// Voice https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Duration uint32 `json:"duration"`
}

// Contact https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"` //nolint:tagliatelle
	FirstName   string `json:"first_name"`   //nolint:tagliatelle
	LastName    string `json:"last_name"`    //nolint:tagliatelle
}

// Location https://core.telegram.org/bots/api#location
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Sticker https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Width    uint32 `json:"width"`
	Height   uint32 `json:"height"`
	FileSize uint32 `json:"file_size"` //nolint:tagliatelle
}

// Animation https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID   string `json:"file_id"` //nolint:tagliatelle
	Width    uint32 `json:"width"`
	Height   uint32 `json:"height"`
	Duration uint32 `json:"duration"`
}

// Story https://core.telegram.org/bots/api#story
type Story struct {
	ID int64 `json:"id"`
}

// Dice https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
}

// Game https://core.telegram.org/bots/api#game
type Game struct {
	Title string `json:"title"`
}

// Invoice https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title string `json:"title"`
}

// Giveaway https://core.telegram.org/bots/api#giveaway
type Giveaway struct {
	Chats []Chat `json:"chats"`
}
