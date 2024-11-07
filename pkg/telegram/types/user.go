package types

// User https://core.telegram.org/bots/api#user
type User struct {
	ID                      int64   `json:"id"`                                    // Unique identifier for this user or bot.
	IsBot                   bool    `json:"is_bot"`                                // True, if this user is a bot.
	FirstName               string  `json:"first_name"`                            // User's or bot's first name.
	LastName                *string `json:"last_name,omitempty"`                   // Optional. User's or bot's last name.
	Username                *string `json:"username,omitempty"`                    // Optional. User's or bot's username.
	LanguageCode            *string `json:"language_code,omitempty"`               // Optional. IETF language tag of the user's language. (Example: ru, en-US)
	IsPremium               *bool   `json:"is_premium,omitempty"`                  // Optional. True, if this user is a Telegram Premium user.
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu,omitempty"`    // Optional. True, if this user added the bot to the attachment menu.
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`             // Optional. True, if the bot can be invited to groups.
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"` // Optional. True, if privacy mode is disabled for the bot.
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`     // Optional. True, if the bot supports inline queries.
	CanConnectToBusiness    *bool   `json:"can_connect_to_business,omitempty"`     // Optional. True, if the bot can connect to a Telegram Business account.
	HasMainWebApp           *bool   `json:"has_main_web_app,omitempty"`            // Optional. True, if the bot has a main Web App.
}

func (u User) LanguageWithDefault() string {
	if u.LanguageCode != nil {
		return *u.LanguageCode
	}
	return "ru"
}
