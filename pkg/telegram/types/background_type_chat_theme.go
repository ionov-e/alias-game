package types

// BackgroundTypeChatTheme
// The background is taken directly from a built-in chat theme.
type BackgroundTypeChatTheme struct {
	// Type of the background, always "chat_theme"
	Type string `json:"type"`
	// ThemeName Name of the chat theme, which is usually an emoji
	ThemeName string `json:"theme_name"`
}

// BackgroundType returns the type of the background, always "chat_theme"
func (b BackgroundTypeChatTheme) BackgroundType() string {
	return b.Type
}
