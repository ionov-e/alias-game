package types

// SuccessfulPayment
// This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double)
	TotalAmount int `json:"total_amount"`
	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`
	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	// Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}
