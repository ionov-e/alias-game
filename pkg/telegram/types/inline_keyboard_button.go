package types

// InlineKeyboardButton represents one button of an inline keyboard. Exactly one of the optional fields must be used to specify the type of button.
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                                       // Label text on the button.
	URL                          *string                      `json:"url,omitempty"`                              // Optional. HTTP or tg:// URL to be opened when the button is pressed.
	CallbackData                 *string                      `json:"callback_data,omitempty"`                    // Optional. Data sent in a callback query to the bot when the button is pressed.
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // Optional. Web App description that launches when pressed. Available only in private chats.
	LoginURL                     *LoginUrl                    `json:"login_url,omitempty"`                        // Optional. HTTPS URL to automatically authorize the user.
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`              // Optional. Prompts the user to select a chat and inserts bot's username and specified query in the input field.
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"` // Optional. Inserts the bot's username and specified query in the current chat's input field.
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // Optional. Prompts the user to select a chat of the specified type, inserts bot's username and specified query in the input field.
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`                        // Optional. Button that copies specified text to the clipboard.
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // Optional. Description of the game that will be launched when pressed.
	Pay                          *bool                        `json:"pay,omitempty"`                              // Optional. Set to true to send a Pay button; replaces "⭐" and "XTR" in text with Telegram Star icon.
}
