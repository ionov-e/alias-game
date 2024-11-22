package types

type PreCheckoutQuery struct {
	// Unique query identifier
	ID string `json:"id"`
	// User who sent the query
	From User `json:"from"`
	// Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Optional. Chosen shipping option identifier
	ShippingOptionID string `json:"shipping_option_id"`
	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}
