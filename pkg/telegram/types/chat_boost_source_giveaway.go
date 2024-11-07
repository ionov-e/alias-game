package types

type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`                     // Source of the boost, always “giveaway”
	GiveawayMessageID int    `json:"giveaway_message_id"`        // Identifier of a message in the chat with the giveaway
	User              User   `json:"user,omitempty"`             // Optional. User that won the prize in the giveaway if any; for Telegram Premium giveaways only
	PrizeStarCount    int    `json:"prize_star_count,omitempty"` // The number of Telegram Stars to be split between giveaway winners
	IsUnclaimed       bool   `json:"is_unclaimed,omitempty"`     // True, if the giveaway was completed, but there was no user to win the prize
}

func (c ChatBoostSourceGiveaway) ChatBoostSource() string { return c.Source }
