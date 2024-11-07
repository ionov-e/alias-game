package types

// GiveawayCreated
// This object represents a service message about the creation of a scheduled giveaway.
type GiveawayCreated struct {
	// PrizeStarCount The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count"`
}
