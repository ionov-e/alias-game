package types

// EncryptedCredentials
// Describes data required for decrypting and authenticating EncryptedPassportElement.
type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data
	Data string `json:"data"`
	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`
	// Base64-encoded secret, encrypted with the bot's public RSA key
	Secret string `json:"secret"`
}
