package status

type Mention struct {
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The account ID of the mentioned user.
	Id string `json:"id"`
	// Type: String
	// Description: The username of the mentioned user.
	Username string `json:"username"`
	// Type: String (URL)
	// Description: The location of the mentioned user’s profile.
	Url string `json:"url"`
	// Type: String
	// Description: The webfinger acct: URI of the mentioned user. Equivalent to username for local users, or username@domain for remote users.
	Acct string `json:"acct"`
}
