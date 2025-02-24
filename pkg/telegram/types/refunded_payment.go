package types

// RefundedPayment
// This object contains basic information about a refunded payment.
type RefundedPayment struct {
	// Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars. Currently, always “XTR”
	Currency string `json:"currency"`
	// Total refunded price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45, total_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	// Optional. Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}
