package types

import (
	"encoding/json"
	"fmt"
)

// SendAnimation represents a request to send an animation file (GIF or H.264/MPEG-4 AVC video). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future. https://core.telegram.org/bots/api#sendanimation
type SendAnimation struct {
	// Optional: Identifier of the business connection for sending the message
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`
	// Optional. Identifier for the target message thread in forums
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// InputFile or String. Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More: https://core.telegram.org/bots/api#sending-files
	Animation InputFileOrString `json:"animation"`
	// Optional: Duration of the animation in seconds
	Duration int `json:"duration,omitempty"`
	// Optional: Animation width
	Width int `json:"width,omitempty"`
	// Optional: Animation height
	Height int `json:"height,omitempty"`
	// Optional: InputFile or String. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More: https://core.telegram.org/bots/api#sending-files
	Thumbnail *InputFileOrString `json:"thumbnail,omitempty"`
	// Optional: Caption for the animation, up to 1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional: Mode for parsing entities in the caption
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional: A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional: Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`
	// Optional: Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
	// Optional: Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional: Protect the message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional: Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional: Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional: Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (s SendAnimation) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendAnimation: %w", err)
	}
	return jsonBytes, nil
}
