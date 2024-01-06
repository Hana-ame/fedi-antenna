package model

// form mastodon note
type LocalNote struct {
	ID string `gorm:"primarykey"`
	// : actor,
	AttributedTo string

	Status      string   `json:"status"`
	MediaIDs    []string `json:"media_ids" gorm:"serializer:json"`
	InReplyToID string   `json:"in_reply_to_id,omitempty"`
	Sensitive   bool     `json:"sensitive"`
	SpoilerText string   `json:"spoiler_text"`
	Visibility  string   `json:"visibility"`

	// dunno
	// ScheduledAt string   `json:"scheduled_at"`

	// utils.Now()
	Published int64

	DeletedAt int64
}
