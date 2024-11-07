package types

// InputPollOption contains information about an answer option to be sent in a poll.
type InputPollOption struct {
	Text          string           `json:"text"`                      // Option text, 1-100 characters.
	TextParseMode *string          `json:"text_parse_mode,omitempty"` // Optional. Mode for parsing entities in the text.
	TextEntities  *[]MessageEntity `json:"text_entities,omitempty"`   // Optional. List of special entities in the option text.
}
