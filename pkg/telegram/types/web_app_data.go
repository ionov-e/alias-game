package types

// WebAppData
// Describes data sent from a Web App to the bot.
type WebAppData struct {
	// Data The data. Be aware that a bad client can send arbitrary data in this field.
	Data *string `json:"data,omitempty"`
	// ButtonText Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
	ButtonText *string `json:"button_text,omitempty"`
}
