package types

type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face for mask placement: "forehead", "eyes", "mouth", or "chin"
	XShift float64 `json:"x_shift"` // X-axis shift by mask width scale, from left to right
	YShift float64 `json:"y_shift"` // Y-axis shift by mask height scale, from top to bottom
	Scale  float64 `json:"scale"`   // Mask scaling coefficient
}
