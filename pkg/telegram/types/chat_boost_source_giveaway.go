package types

type ChatBoostSourceGiveaway struct {
	// Source of the boost, always “giveaway”
	Source string `json:"source"`
	// Identifier of a message in the chat with the giveaway
	GiveawayMessageID int `json:"giveaway_message_id"`
	// Optional. User that won the prize in the giveaway if any; for Telegram Premium giveaways only
	User *User `json:"user,omitempty"`
	// Optional. The number of Telegram Stars to be split between giveaway winners
	PrizeStarCount int `json:"prize_star_count,omitempty"`
	// Optional. True, if the giveaway was completed, but there was no user to win the prize
	IsUnclaimed bool `json:"is_unclaimed,omitempty"`
}

func (c ChatBoostSourceGiveaway) ChatBoostSource() string { return c.Source }
