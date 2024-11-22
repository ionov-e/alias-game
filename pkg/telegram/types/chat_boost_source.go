package types

import "fmt"

type ChatBoostSource interface {
	ChatBoostSource() string
}

func ChatBoostSourceFactory(source string, data map[string]interface{}) (ChatBoostSource, error) {
	user, userIsPresent := data["user"].(User)

	switch source {
	case "premium":
		if !userIsPresent {
			return nil, fmt.Errorf("invalid or missing 'user' field for 'premium'")
		}
		return ChatBoostSourcePremium{
			Source: "premium",
			User:   user,
		}, nil
	case "gift_code":
		if !userIsPresent {
			return nil, fmt.Errorf("invalid or missing 'user' field for 'gift_code'")
		}
		return ChatBoostSourceGiftCode{
			Source: "gift_code",
			User:   user,
		}, nil
	case "giveaway":
		giveawayMessageID, ok := data["giveaway_message_id"].(int)
		if !ok {
			return nil, fmt.Errorf("invalid or missing 'giveaway_message_id' field for 'giveaway'")
		}

		chatBoostSourceGiveaway := ChatBoostSourceGiveaway{
			Source:            "giveaway",
			GiveawayMessageID: giveawayMessageID,
		}

		if userIsPresent {
			chatBoostSourceGiveaway.User = &user
		}

		if prizeStarCount, ok := data["prize_star_count"].(int); ok {
			chatBoostSourceGiveaway.PrizeStarCount = prizeStarCount
		}

		if isUnclaimed, ok := data["is_unclaimed"].(bool); ok {
			chatBoostSourceGiveaway.IsUnclaimed = isUnclaimed
		}

		return chatBoostSourceGiveaway, nil
	default:
		return nil, fmt.Errorf("unknown ChatBoostSource type: %s", source)
	}
}
