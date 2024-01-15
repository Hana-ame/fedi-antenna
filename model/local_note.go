package model

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

	DeletedAt int64
}
