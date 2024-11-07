package types

import "fmt"

type ChatBoostSource interface {
	ChatBoostSource() string
}

func ChatBoostSourceFactory(source string, data map[string]interface{}) (ChatBoostSource, error) {
	switch source {
	case "premium":
		return ChatBoostSourcePremium{
			Source: "premium",
			User:   data["user"].(User),
		}, nil
	case "gift_code":
		return ChatBoostSourceGiftCode{
			Source: "gift_code",
			User:   data["user"].(User),
		}, nil
	case "giveaway":
		return ChatBoostSourceGiveaway{
			Source:            "giveaway",
			GiveawayMessageID: data["giveaway_message_id"].(int),
			User:              data["user"].(User),            // May be empty
			PrizeStarCount:    data["prize_star_count"].(int), // Optional field
			IsUnclaimed:       data["is_unclaimed"].(bool),    // Optional field
		}, nil
	default:
		return nil, fmt.Errorf("unknown ChatBoostSource type: %s", source)
	}
}
