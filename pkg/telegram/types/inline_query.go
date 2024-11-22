package types

type InlineQuery struct {
	ID       string    `json:"id"`                  // Unique identifier for this query
	From     User      `json:"from"`                // Sender
	Query    string    `json:"query"`               // Text of the query (up to 256 characters)
	Offset   string    `json:"offset"`              // Offset of the results to be returned, can be controlled by the bot
	ChatType string    `json:"chat_type,omitempty"` // Optional. Type of the chat from which the inline query was sent
	Location *Location `json:"location,omitempty"`  // Optional. Sender location, only for bots that request user location
}
