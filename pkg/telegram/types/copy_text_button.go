package types

// CopyTextButton represents an inline keyboard button that copies specified text to the clipboard.
type CopyTextButton struct {
	Text string `json:"text"` // Text to be copied to the clipboard; 1-256 characters.
}
