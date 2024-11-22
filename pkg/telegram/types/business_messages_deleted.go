package types

type BusinessMessagesDeleted struct {
	// Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id"`
	// Information about a chat in the business account
	Chat Chat `json:"chat"`
	//nolint:tagliatelle    // List of identifiers of deleted messages in the business account chat
	MessageIDs []int `json:"message_ids"`
}
