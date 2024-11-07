package types

import "fmt"

// PaidMedia
// Interface for different types of paid media.
type PaidMedia interface {
	// PaidMediaType returns the type of the paid media
	PaidMediaType() string
}

// PaidMediaFactory
// Factory function to create PaidMedia implementations based on media type.
func PaidMediaFactory(mediaType string, data map[string]interface{}) (PaidMedia, error) {
	switch mediaType {
	case "preview":
		return PaidMediaPreview{
			Type:     "preview",
			Width:    intPtr(data["width"].(int)),
			Height:   intPtr(data["height"].(int)),
			Duration: intPtr(data["duration"].(int)),
		}, nil
	case "photo":
		return PaidMediaPhoto{
			Type:  "photo",
			Photo: data["photo"].([]PhotoSize),
		}, nil
	case "video":
		return PaidMediaVideo{
			Type:  "video",
			Video: data["video"].(Video),
		}, nil
	default:
		return nil, fmt.Errorf("unknown PaidMedia type: %s", mediaType)
	}
}

func intPtr(i int) *int {
	return &i
}
