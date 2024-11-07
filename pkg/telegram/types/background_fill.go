package types

import (
	"errors"
	"fmt"
)

// BackgroundFill
// Interface for different types of background fills.
type BackgroundFill interface {
	BackgroundType() string
}

// BackgroundFillFactory creates a BackgroundFill object based on the provided type and data.
func BackgroundFillFactory(data map[string]interface{}) (BackgroundFill, error) {
	bgType, ok := data["type"].(string)
	if !ok {
		return nil, errors.New("type field is required")
	}

	switch bgType {
	case "freeform_gradient":
		colors, ok := data["colors"].([]int)
		if !ok {
			return nil, fmt.Errorf("invalid colors data for BackgroundFillFreeformGradient")
		}
		return BackgroundFillFreeformGradient{
			Type:   bgType,
			Colors: colors,
		}, nil

	case "gradient":
		topColor, ok1 := data["top_color"].(int)
		bottomColor, ok2 := data["bottom_color"].(int)
		rotationAngle, ok3 := data["rotation_angle"].(int)
		if !ok1 || !ok2 || !ok3 {
			return nil, fmt.Errorf("invalid gradient data for BackgroundFillGradient")
		}
		return BackgroundFillGradient{
			Type:          bgType,
			TopColor:      topColor,
			BottomColor:   bottomColor,
			RotationAngle: rotationAngle,
		}, nil

	case "solid":
		color, ok := data["color"].(int)
		if !ok {
			return nil, fmt.Errorf("invalid color data for BackgroundFillSolid")
		}
		return BackgroundFillSolid{
			Type:  bgType,
			Color: color,
		}, nil

	default:
		return nil, fmt.Errorf("unknown BackgroundFill type: %s", bgType)
	}
}
