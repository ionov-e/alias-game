package types

// User https://core.telegram.org/bots/api#user
type User struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`     //nolint:tagliatelle
	FirstName string `json:"first_name"` //nolint:tagliatelle
	// Optional: IETF language tag
	Language string `json:"language_code"` //nolint:tagliatelle
}

func (u User) LanguageWithDefault() string {
	if u.Language != "" {
		return u.Language
	}
	return "ru" // TODO check if correct string
}
