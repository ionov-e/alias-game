package types

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`                          // Invite link URL.
	Creator                 User   `json:"creator"`                              // Creator of the link.
	CreatesJoinRequest      bool   `json:"creates_join_request"`                 // True if joining requires admin approval.
	IsPrimary               bool   `json:"is_primary"`                           // True if the link is primary.
	IsRevoked               bool   `json:"is_revoked"`                           // True if the link is revoked.
	Name                    string `json:"name,omitempty"`                       // Optional. Invite link name.
	ExpireDate              int64  `json:"expire_date,omitempty"`                // Optional. Expiration date (Unix timestamp).
	MemberLimit             int    `json:"member_limit,omitempty"`               // Optional. Max users that can join using this link.
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"` // Optional. Pending join request count.
	SubscriptionPeriod      int    `json:"subscription_period,omitempty"`        // Optional. Active subscription duration in seconds.
	SubscriptionPrice       int    `json:"subscription_price,omitempty"`         // Optional. Cost in Telegram Stars for chat membership.
}
