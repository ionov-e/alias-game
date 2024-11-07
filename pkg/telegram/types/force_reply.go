package types

// ForceReply represents a request for Telegram clients to display a reply interface.
type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`                       // Shows reply interface to the user.
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"` // Optional. Placeholder shown in the input field.
	Selective             *bool   `json:"selective,omitempty"`               // Optional. Forces reply from specific users only.
}
