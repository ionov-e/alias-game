package types

import "fmt"

// MessageOrigin https://core.telegram.org/bots/api#messageorigin
type MessageOrigin interface {
	MessageOriginType() string
	MessageOriginDate() int
}

func MessageOriginFactory(typeValue string, data map[string]interface{}) (MessageOrigin, error) {
	switch typeValue {
	case "user":
		return MessageOriginUser{
			Type:       "user",
			Date:       data["date"].(int),
			SenderUser: data["sender_user"].(User),
		}, nil
	case "hidden_user":
		return MessageOriginHiddenUser{
			Type:           "hidden_user",
			Date:           data["date"].(int),
			SenderUserName: data["sender_user_name"].(string),
		}, nil
	case "chat":
		authorSignature, ok := data["author_signature"].(string)
		if !ok {
			return MessageOriginChat{
				Type:       "chat",
				Date:       data["date"].(int),
				SenderChat: data["sender_chat"].(Chat),
			}, nil
		}
		return MessageOriginChat{
			Type:            "chat",
			Date:            data["date"].(int),
			SenderChat:      data["sender_chat"].(Chat),
			AuthorSignature: authorSignature,
		}, nil
	case "channel":
		authorSignature, ok := data["author_signature"].(string)
		if !ok {
			return MessageOriginChannel{
				Type:      "channel",
				Date:      data["date"].(int),
				Chat:      data["chat"].(Chat),
				MessageID: data["message_id"].(int),
			}, nil
		}
		return MessageOriginChannel{
			Type:            "channel",
			Date:            data["date"].(int),
			Chat:            data["chat"].(Chat),
			MessageID:       data["message_id"].(int),
			AuthorSignature: authorSignature,
		}, nil
	default:
		return nil, fmt.Errorf("unknown MessageOrigin type: %s", typeValue)
	}
}
