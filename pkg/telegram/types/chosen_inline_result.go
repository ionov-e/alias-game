package types

type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`                   // The unique identifier for the result that was chosen
	From            User      `json:"from"`                        // The user that chose the result
	Location        *Location `json:"location,omitempty"`          // Optional. Sender location, only for bots that require user location
	InlineMessageID *string   `json:"inline_message_id,omitempty"` // Optional. Identifier of the sent inline message
	Query           string    `json:"query"`                       // The query that was used to obtain the result
}
