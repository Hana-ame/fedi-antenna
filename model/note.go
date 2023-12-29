package model

import "github.com/Hana-ame/fedi-antenna/core/utils"

// order from mastodon
type Note struct {
	// not used?
	Context any `json:"omitempty" gorm:"-"`

	// /users/:name/statues/:id
	ID string `json:"id" gorm:"primarykey"`
	// alwasy "Note"
	Type string `json:"type" gorm:"-"`

	// this should be html marshaled
	Summary *string `json:"summary"`

	InReplyTo *string `json:"inReplyTo"`

	Published string `json:"published"`

	// this one is for browser
	URL string `json:"url"`

	// user's activitypub ID / url
	AttributeTo string `json:"attributeTo"`

	To []string `json:"to" gorm:"serializer:json"`

	Cc []string `json:"cc" gorm:"serializer:json"`

	Sensitive bool `json:"sensitive"`

	// same to ID
	AtomUri string `json:"atomUri" gorm:"-"`
	// same to inReplyTo
	InReplyToAtomUri *string `json:"inReplyToAtomUri" gorm:"-"`

	// not implemented
	// Conversation string `json:"conversation"`

	// should be html marshaled
	Content string `json:"content"`

	// not implemented
	// ContentMap any `json:"contentMap"`

	// attachments, not implicated
	Attachment []any `json:"attachment" gorm:"-"`

	// if you @someone
	Tag []*Mention `json:"tag" gorm:"serializer:json"`

	Replies *Collection `json:"replies" gorm:"foreignKey:ID;references:ID"`
}

var NoteContext = []any{
	"https://www.w3.org/ns/activitystreams",
	utils.NewMapFromKV([]*utils.KV{
		{Key: "ostatus", Value: "http://ostatus.org#"},
		{Key: "atomUri", Value: "ostatus:atomUri"},
		{Key: "inReplyToAtomUri", Value: "ostatus:inReplyToAtomUri"},
		{Key: "conversation", Value: "ostatus:conversation"},
		{Key: "sensitive", Value: "as:sensitive"},
		{Key: "toot", Value: "http://joinmastodon.org/ns#"},
		{Key: "votersCount", Value: "toot:votersCount"},
	}),
}

func (o *Note) Autofill() {
	o.Context = NoteContext
	o.ID = "Note"
	o.AtomUri = o.ID
	o.InReplyToAtomUri = o.InReplyTo
}
