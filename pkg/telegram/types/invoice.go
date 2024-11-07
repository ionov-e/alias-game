package types

// Invoice https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title          string `json:"title"`           // Product name.
	Description    string `json:"description"`     // Product description.
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter to generate this invoice.
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code, or "XTR" for Telegram Stars.
	TotalAmount    int    `json:"total_amount"`    // Total price in smallest units of the currency (e.g., 145 for US$1.45).
}
