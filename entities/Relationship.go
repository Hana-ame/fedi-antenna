package entities

type Relationship struct {
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The account ID.
	Attributes string `json:"Attributes"`
	// Type: Boolean
	// Description: Are you following this user?
	Following bool `json:"following"`
	// Type: Boolean
	// Description: Are you receiving this user’s boosts in your home timeline?
	ShowingReblogs bool `json:"showing_reblogs"`
	// Type: Boolean
	// Description: Have you enabled notifications for this user?
	Notifying bool `json:"notifying"`
	// Type: Array of String (ISO 639-1 language two-letter code)
	// Description: Which languages are you following from this user?
	Languages []string `json:"languages"`
	// Type: Boolean
	// Description: Are you followed by this user?
	FollowedBy bool `json:"followed_by"`
	// Type: Boolean
	// Description: Are you blocking this user?
	Blocking bool `json:"blocking"`
	// Type: Boolean
	// Description: Is this user blocking you?
	BlockedBy bool `json:"blocked_by"`
	// Type: Boolean
	// Description: Are you muting this user?
	Muting bool `json:"muting"`
	// Type: Boolean
	// Description: Are you muting notifications from this user?
	MutingNotifications bool `json:"muting_notifications"`
	// Type: Boolean
	// Description: Do you have a pending follow request for this user?
	Requested bool `json:"requested"`
	// Type: Boolean
	// Description: Has this user requested to follow you?
	RequestedBy bool `json:"requested_by"`
	// Type: Boolean
	// Description: Are you blocking this user’s domain?
	DomainBlocking bool `json:"domain_blocking"`
	// Type: Boolean
	// Description: Are you featuring this user on your profile?
	Endorsed bool `json:"endorsed"`
	// Type: String
	// Description: This user’s profile bio
	Note string `json:"note"`
}
