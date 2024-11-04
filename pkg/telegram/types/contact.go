package types

// Contact https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"` //nolint:tagliatelle
	FirstName   string `json:"first_name"`   //nolint:tagliatelle
	LastName    string `json:"last_name"`    //nolint:tagliatelle
}
