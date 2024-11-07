package types

// EncryptedPassportElement
// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	// Element type. One of the predefined types
	Type string `json:"type"`
	// Base64-encoded encrypted Telegram Passport element data
	Data *string `json:"data,omitempty"`
	// User's verified phone number
	PhoneNumber *string `json:"phone_number,omitempty"`
	// User's verified email address
	Email *string `json:"email,omitempty"`
	// Array of encrypted files with documents provided by the user
	Files []PassportFile `json:"files,omitempty"`
	// Encrypted file with the front side of the document
	FrontSide *PassportFile `json:"front_side,omitempty"`
	// Encrypted file with the reverse side of the document
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`
	// Encrypted file with the selfie of the user holding a document
	Selfie *PassportFile `json:"selfie,omitempty"`
	// Array of encrypted files with translated versions of documents
	Translation []PassportFile `json:"translation,omitempty"`
	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash"`
}
