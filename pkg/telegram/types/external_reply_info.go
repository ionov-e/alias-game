package types

// ExternalReplyInfo https://core.telegram.org/bots/api#externalreplyinfo
type ExternalReplyInfo struct {
	MessageOrigin      MessageOrigin       `json:"origin"`                         // Origin of the message being replied to
	Chat               *Chat               `json:"chat,omitempty"`                 // Optional: Chat of the original message (if supergroup or channel)
	MessageID          int                 `json:"message_id,omitempty"`           // Optional: Message ID in the original chat
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Optional: Options for link preview
	Animation          *Animation          `json:"animation,omitempty"`            // Optional: Animation info, if message is an animation
	Audio              *Audio              `json:"audio,omitempty"`                // Optional: Audio file info, if message is audio
	Document           *Document           `json:"document,omitempty"`             // Optional: General file info, if message is a document
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`           // Optional: Paid media info, if message contains paid media
	Photo              []PhotoSize         `json:"photo,omitempty"`                // Optional: Available sizes, if message is a photo
	Sticker            *Sticker            `json:"sticker,omitempty"`              // Optional: Sticker info, if message is a sticker
	Story              *Story              `json:"story,omitempty"`                // Optional: Story info, if message is a forwarded story
	Video              *Video              `json:"video,omitempty"`                // Optional: Video info, if message is a video
	VideoNote          *VideoNote          `json:"video_note,omitempty"`           // Optional: Video note info, if message is a video note
	Voice              *Voice              `json:"voice,omitempty"`                // Optional: Voice message info
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`    // Optional: True if media is covered by spoiler animation
	Contact            *Contact            `json:"contact,omitempty"`              // Optional: Contact info, if message is a shared contact
	Dice               *Dice               `json:"dice,omitempty"`                 // Optional: Dice with random value
	Game               *Game               `json:"game,omitempty"`                 // Optional: Game info, if message is a game
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`             // Optional: Scheduled giveaway info
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`     // Optional: Completed giveaway with winners
	Invoice            *Invoice            `json:"invoice,omitempty"`              // Optional: Invoice info, if message is an invoice
	Location           *Location           `json:"location,omitempty"`             // Optional: Location info, if message is a shared location
	Poll               *Poll               `json:"poll,omitempty"`                 // Optional: Poll info, if message is a native poll
	Venue              *Venue              `json:"venue,omitempty"`                // Optional: Venue info, if message is a venue
}
