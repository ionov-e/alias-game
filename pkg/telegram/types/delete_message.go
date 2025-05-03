package types

import (
	"encoding/json"
	"fmt"
)

// DeleteMessage
// Use this method to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
// Returns True on success.
// https://core.telegram.org/bots/api#deletemessage
type DeleteMessage struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID int64 `json:"chat_id"` // Integer or String
	// Identifier of the message to delete
	MessageID int64 `json:"message_id"`
}

func (s DeleteMessage) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling DeleteMessage: %w", err)
	}
	return jsonBytes, nil
}
