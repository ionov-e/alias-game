package types

// PassportData
// Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Array with information about documents and other Telegram Passport elements that was shared with the bot
	Data []EncryptedPassportElement `json:"data"`
	// Encrypted credentials required to decrypt the data
	Credentials EncryptedCredentials `json:"credentials"`
}
