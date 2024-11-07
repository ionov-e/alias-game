package types

import "fmt"

type ReactionType interface {
	ReactionType() string
}

func ReactionTypeFactory(typeValue string, data map[string]interface{}) (ReactionType, error) {
	switch typeValue {
	case "custom_emoji":
		return ReactionTypeCustomEmoji{
			Type:          "custom_emoji",
			CustomEmojiID: data["custom_emoji_id"].(string),
		}, nil
	case "emoji":
		return ReactionTypeEmoji{
			Type:  "emoji",
			Emoji: data["emoji"].(string),
		}, nil
	case "paid":
		return ReactionTypePaid{
			Type: "paid",
		}, nil
	default:
		return nil, fmt.Errorf("unknown ReactionType: %s", typeValue)
	}
}
