package types

// GiveawayCompleted
// This object represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	// WinnerCount Number of winners in the giveaway
	WinnerCount int `json:"winner_count"`
	// UnclaimedPrizeCount Number of undistributed prizes
	UnclaimedPrizeCount int `json:"unclaimed_prize_count"`
	// GiveawayMessage Message with the giveaway that was completed, if it wasn't deleted
	GiveawayMessage Message `json:"giveaway_message"`
	// IsStarGiveaway True, if the giveaway is a Telegram Star giveaway
	IsStarGiveaway bool `json:"is_star_giveaway"`
}
