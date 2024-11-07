package types

type TextQuote struct {
	Text     string          `json:"text"`                // The quoted text
	Entities []MessageEntity `json:"entities,omitempty"`  // Optional: Special entities in the quote
	Position int             `json:"position"`            // Position in UTF-16 code units
	IsManual bool            `json:"is_manual,omitempty"` // Optional: True if chosen manually
}
