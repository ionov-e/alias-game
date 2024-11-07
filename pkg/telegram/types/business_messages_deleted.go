package types

type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"` // Unique identifier of the business connection
	Chat                 Chat   `json:"chat"`                   // Information about a chat in the business account
	MessageIDs           []int  `json:"message_ids"`            // List of identifiers of deleted messages in the business account chat
}
