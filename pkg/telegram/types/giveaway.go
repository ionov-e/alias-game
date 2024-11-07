package types

// Giveaway
// This object represents a message about a scheduled giveaway.
type Giveaway struct {
	// Chats The list of chats which the user must join to participate in the giveaway
	Chats []Chat `json:"chats"`
	// WinnersSelectionDate Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnersSelectionDate int `json:"winners_selection_date"`
	// WinnerCount The number of users which are supposed to be selected as winners of the giveaway
	WinnerCount int `json:"winner_count"`
	// OnlyNewMembers True, if only users who join the chats after the giveaway started should be eligible to win
	OnlyNewMembers bool `json:"only_new_members"`
	// HasPublicWinners True, if the list of giveaway winners will be visible to everyone
	HasPublicWinners bool `json:"has_public_winners"`
	// PrizeDescription Description of additional giveaway prize
	PrizeDescription string `json:"prize_description"`
	// CountryCodes A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come
	CountryCodes []string `json:"country_codes"`
	// PrizeStarCount The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count"`
	// PremiumSubscriptionMonthCount The number of months the Telegram Premium subscription won from the giveaway will be active for
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count"`
}
