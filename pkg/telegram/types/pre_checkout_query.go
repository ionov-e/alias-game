package types

type PreCheckoutQuery struct {
	ID               string     `json:"id"`                   // Unique query identifier
	From             User       `json:"from"`                 // User who sent the query
	Currency         string     `json:"currency"`             // Three-letter ISO 4217 currency code
	TotalAmount      int        `json:"total_amount"`         // Total price in the smallest units of the currency
	InvoicePayload   string     `json:"invoice_payload"`      // Bot-specified invoice payload
	ShippingOptionID string     `json:"shipping_option_id"`   // Optional. Chosen shipping option identifier
	OrderInfo        *OrderInfo `json:"order_info,omitempty"` // Optional. Order information provided by the user
}
