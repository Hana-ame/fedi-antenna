package model

import "github.com/Hana-ame/fedi-antenna/mastodon/entities/poll"

// 用来保存note的结构
// 转mastodon格式，在mastodon访问时
// 转activitypub格式，在发送时
// 从mastodon格式，在收到人工发送时
// 从activitypub格式，在inbox受到或者准备从远处拉取时
// form mastodon note
type LocalNote struct {
	// "https://" + host + "/users/" + name + "/statuses/" + id
	ID string `gorm:"primarykey"`
	// actor "https://" + host + "/users/" + name
	AttributedTo string

	// text, html encoded
	Status      string   `json:"status"`
	MediaIDs    []string `json:"media_ids" gorm:"serializer:json"`
	InReplyToID string   `json:"in_reply_to_id,omitempty"`
	Sensitive   bool     `json:"sensitive"`
	SpoilerText string   `json:"spoiler_text"`
	// {public, unlisted, private, direct}
	Visibility string `json:"visibility"`

	ReblogsCount    int `json:"reblogs_count"`
	FavouritesCount int `json:"favourites_count"`
	RepliesCount    int

	// dunno
	// ScheduledAt string   `json:"scheduled_at"`
	InReplyToAccountId string

	// utils.Now()
	// also the ID for mastodon strconv.Itoa()
	Published int64

	Poll *LocalPoll `gorm:"foreignKey:Published;references:ID"`

	DeletedAt int64
}

type LocalPoll struct {
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The ID of the poll in the database.
	ID int64 `gorm:"primarykey"`
	// Type: NULLABLE String (ISO 8601 Datetime), or null if the poll does not end
	// Description: When the poll ends.
	ExpiresAt *int64 `json:"expires_at"`
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
	Options []poll.Option `json:"options" gorm:"serializer:json"`
	// Type: Array of CustomEmoji
	// Description: Custom emoji to be used for rendering poll options.
	// Emojis []CustomEmoji `json:"emojis"`
	// Type: Boolean
	// Description: When called with a user token, has the authorized user voted?
	// Voted bool `json:"voted,omitempty"`
	// Type: Array of Integer
	// Description: When called with a user token, which options has the authorized user chosen? Contains an array of index values for options.
	// OwnVotes []int `json:"own_votes,omitempty"`
}
