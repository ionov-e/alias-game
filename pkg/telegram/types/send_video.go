package types

import (
	"encoding/json"
	"fmt"
)

// SendVideo represents a request to send a video file.
type SendVideo struct {
	// Optional: Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Optional: Identifier for the target message thread in forums
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data.
	Video InputFileOrString `json:"video"`
	// Optional: Duration of the video in seconds
	Duration int `json:"duration,omitempty"`
	// Optional: Video width
	Width int `json:"width,omitempty"`
	// Optional: Video height
	Height int `json:"height,omitempty"`
	// Optional: Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More:https://core.telegram.org/bots/api#sending-files
	Thumbnail *InputFileOrString `json:"thumbnail,omitempty"`
	// Optional: Caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional: Mode for parsing entities in the video caption. See formatting options: https://core.telegram.org/bots/api#formatting-options
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional: List of special entities in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional: Display caption above the media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`
	// Optional: Display the video with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
	// Optional: Indicate if the video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
	// Optional: Send the message without sound
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional: Protect the message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional: Effect ID for adding an effect to the message
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional: Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (s SendVideo) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendVideo: %w", err)
	}
	return jsonBytes, nil
}
