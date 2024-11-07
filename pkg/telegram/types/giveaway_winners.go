package types

// GiveawayWinners
// This object represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	// Chat The chat that created the giveaway
	Chat Chat `json:"chat"`
	// GiveawayMessageID Identifier of the message with the giveaway in the chat
	GiveawayMessageID int `json:"giveaway_message_id"`
	// WinnersSelectionDate Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int `json:"winners_selection_date"`
	// WinnerCount Total number of winners in the giveaway
	WinnerCount int `json:"winner_count"`
	// Winners List of up to 100 winners of the giveaway
	Winners []User `json:"winners"`
	// AdditionalChatCount The number of other chats the user had to join in order to be eligible for the giveaway
	AdditionalChatCount int `json:"additional_chat_count"`
	// PrizeStarCount The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count"`
	// PremiumSubscriptionMonthCount The number of months the Telegram Premium subscription won from the giveaway will be active for
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count"`
	// UnclaimedPrizeCount Number of undistributed prizes
	UnclaimedPrizeCount int `json:"unclaimed_prize_count"`
	// OnlyNewMembers True, if only users who had joined the chats after the giveaway started were eligible to win
	OnlyNewMembers bool `json:"only_new_members"`
	// WasRefunded True, if the giveaway was canceled because the payment for it was refunded
	WasRefunded bool `json:"was_refunded"`
	// PrizeDescription Description of additional giveaway prize
	PrizeDescription string `json:"prize_description"`
}
