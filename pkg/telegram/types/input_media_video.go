package types

// InputMediaVideo represents a video to be sent.
type InputMediaVideo struct {
	// Type of the result, must be video
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More: https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
	// Optional. InputFile (must be posted using multipart/form-data in the usual way that files are uploaded via the browser) or String. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files
	Thumbnail *InputFileOrString `json:"thumbnail,omitempty"`
	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the video caption. See more: https://core.telegram.org/bots/api#formatting-options
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`
	// Optional: Video width
	Width int `json:"width,omitempty"`
	// Optional: Video height
	Height int `json:"height,omitempty"`
	// Optional: Video duration in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Pass True if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
	// Optional. Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// InputMediaType returns the type of the media, which is "video".
func (v *InputMediaVideo) InputMediaType() string {
	return v.Type
}
