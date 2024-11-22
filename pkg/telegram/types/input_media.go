package types

import "fmt"

// InputMedia represents the content of a media message to be sent.
// It includes the InputMediaType method, which returns the type of the media.
type InputMedia interface {
	InputMediaType() string
}

func InputMediaFactory(mediaType string, data map[string]interface{}) (InputMedia, error) {
	switch mediaType {
	case "photo":
		return &InputMediaPhoto{
			Type:                  "photo",
			Media:                 data["media"].(string),
			Caption:               data["caption"].(string),
			ParseMode:             data["parse_mode"].(string),
			CaptionEntities:       data["caption_entities"].([]MessageEntity),
			ShowCaptionAboveMedia: data["show_caption_above_media"].(bool),
			HasSpoiler:            data["has_spoiler"].(bool),
		}, nil
	case "video":
		return &InputMediaVideo{
			Type:                  "video",
			Media:                 data["media"].(string),
			Thumbnail:             data["thumbnail"].(*InputFileOrString),
			Caption:               data["caption"].(string),
			ParseMode:             data["parse_mode"].(string),
			CaptionEntities:       data["caption_entities"].([]MessageEntity),
			ShowCaptionAboveMedia: data["show_caption_above_media"].(bool),
			Width:                 data["width"].(int),
			Height:                data["height"].(int),
			Duration:              data["duration"].(int),
			SupportsStreaming:     data["supports_streaming"].(bool),
			HasSpoiler:            data["has_spoiler"].(bool),
		}, nil
	case "animation":
		return &InputMediaAnimation{
			Type:                  "animation",
			Media:                 data["media"].(string),
			Thumbnail:             data["thumbnail"].(*InputFileOrString),
			Caption:               data["caption"].(string),
			ParseMode:             data["parse_mode"].(string),
			CaptionEntities:       data["caption_entities"].([]MessageEntity),
			ShowCaptionAboveMedia: data["show_caption_above_media"].(bool),
			Width:                 data["width"].(int),
			Height:                data["height"].(int),
			Duration:              data["duration"].(int),
			HasSpoiler:            data["has_spoiler"].(bool),
		}, nil
	case "audio":
		return &InputMediaAudio{
			Type:            "audio",
			Media:           data["media"].(string),
			Thumbnail:       data["thumbnail"].(*InputFileOrString),
			Caption:         data["caption"].(string),
			ParseMode:       data["parse_mode"].(string),
			CaptionEntities: data["caption_entities"].([]MessageEntity),
			Duration:        data["duration"].(int),
			Performer:       data["performer"].(string),
			Title:           data["title"].(string),
		}, nil
	case "document":
		return &InputMediaDocument{
			Type:                        "document",
			Media:                       data["media"].(string),
			Thumbnail:                   data["thumbnail"].(*InputFileOrString),
			Caption:                     data["caption"].(string),
			ParseMode:                   data["parse_mode"].(string),
			CaptionEntities:             data["caption_entities"].([]MessageEntity),
			DisableContentTypeDetection: data["disable_content_type_detection"].(bool),
		}, nil
	default:
		return nil, fmt.Errorf("unknown InputMedia type: %s", mediaType)
	}
}
