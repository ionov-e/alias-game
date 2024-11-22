package types

import (
	"encoding/json"
	"fmt"
)

// SendAudio to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format.
// On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
type SendAudio struct {
	// Optional: Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`
	// Optional: Identifier for the target message thread in forums
	MessageThreadID int `json:"message_thread_id,omitempty"`
	// InputFile or String. Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data
	Audio InputFileOrString `json:"audio"`
	// Optional: Caption for the audio, up to 1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional: Mode for parsing entities in the caption
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional: List of special entities in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional: Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`
	// Optional: Performer of the audio
	Performer string `json:"performer,omitempty"`
	// Optional: Track name
	Title string `json:"title,omitempty"`
	// Optional: Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Thumbnail *InputFileOrString `json:"thumbnail,omitempty"`
	// Optional: Send the message without sound
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional: Protect the message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional: Allow up to 1000 messages per second with paid broadcasting
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`
	// Optional: Effect ID for adding an effect to the message
	MessageEffectID string `json:"message_effect_id,omitempty"`
	// Optional: Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	// Optional: InlineKeyboardMarkup, etc.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (s SendAudio) Bytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return jsonBytes, fmt.Errorf("error marshalling SendAudio: %w", err)
	}
	return jsonBytes, nil
}
