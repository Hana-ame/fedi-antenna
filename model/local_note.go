package model

// form mastodon note
type LocalNote struct {
	// ID is a integer like which means timestamp
	ID string `gorm:"primarykey"`
	// actor
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
	Published int64

	DeletedAt int64
}
