package types

// EncryptedPassportElement
// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	// Element type. One of the predefined types
	Type string `json:"type"`
	// Optional. Base64-encoded encrypted Telegram Passport element data provided by the user; available only for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	Data string `json:"data,omitempty"`
	// Optional. User's verified phone number; available only for “phone_number” type
	PhoneNumber string `json:"phone_number,omitempty"`
	// Optional. User's verified email address; available only for “email” type
	Email string `json:"email,omitempty"`
	// Optional. Array of encrypted files with documents provided by the user; available only for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Files []PassportFile `json:"files,omitempty"`
	// Optional. Encrypted file with the front side of the document, provided by the user; available only for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side,omitempty"`
	// Optional. Encrypted file with the reverse side of the document, provided by the user; available only for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`
	// Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available if requested for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie,omitempty"`
	// Optional. Array of encrypted files with translated versions of documents provided by the user; available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []PassportFile `json:"translation,omitempty"`
	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash"`
}
