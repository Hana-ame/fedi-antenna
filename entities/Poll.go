package entities

import "github.com/Hana-ame/fedi-antenna/mastodon/entities/poll"

type Poll struct {
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The ID of the poll in the database.
	Id string `json:"id"`
	// Type: NULLABLE String (ISO 8601 Datetime), or null if the poll does not end
	// Description: When the poll ends.
	ExpiresAt *string `json:"expires_at"`
	// Type: Boolean
	// Description: Is the poll currently expired?
	Expired bool `json:"expired"`
	// Type: Boolean
	// Description: Does the poll allow multiple-choice answers?
	Multiple bool `json:"multiple"`
	// Type: Integer
	// Description: How many votes have been received.
	VotesCount int `json:"votes_count"`
	// Type: NULLABLE Integer, or null if multiple is false.
	// Description: How many unique accounts have voted on a multiple-choice poll.
	VotersCount *int `json:"voters_count"`
	// Type: Array of Poll::Option
	// Description: Possible answers for the poll.
	Options []poll.Option `json:"options"`
	// Type: Array of CustomEmoji
	// Description: Custom emoji to be used for rendering poll options.
	Emojis []CustomEmoji `json:"emojis"`
	// Type: Boolean
	// Description: When called with a user token, has the authorized user voted?
	Voted bool `json:"voted,omitempty"`
	// Type: Array of Integer
	// Description: When called with a user token, which options has the authorized user chosen? Contains an array of index values for options.
	OwnVotes []int `json:"own_votes,omitempty"`
}
