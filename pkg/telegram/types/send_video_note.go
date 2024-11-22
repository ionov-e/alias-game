package types

import (
	"encoding/json"
	"fmt"
)

// SendVideoNote As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNote struct {
	// Optional: Identifier of the business connection for sending the message
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Optional: Identifier for the target message thread in forums
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. Sending video notes by a URL is currently unsupported. More information: https://core.telegram.org/bots/api#sending-files
	VideoNote InputFileOrString `json:"video_note"`
	// Optional: Duration of the video in seconds
	Duration int `json:"duration,omitempty"`
	// Optional: Diameter of the video message (width & height)
	Length int `json:"length,omitempty"`
	// Optional: Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information: https://core.telegram.org/bots/api#sending-files
	Thumbnail *InputFileOrString `json:"thumbnail,omitempty"`
	// Optional: Send the message silently without notification sound
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional: Prevent the message from being forwarded or saved
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional: Allow paid broadcasting
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional: Message effect ID for adding effects to the message
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional: Information about the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (s SendVideoNote) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendVideoNote: %w", err)
	}
	return jsonBytes, nil
}
