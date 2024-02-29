package model

import (
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// order from mastodon
type Note struct {
	// not used?
	Context any `json:"@context,omitempty" gorm:"-"`

	// /users/:name/statues/:id
	ID string `json:"id" gorm:"primarykey"`

	// always "Note"
	Type string `json:"type" gorm:"-"`

	// this should be html marshaled
	Summary *string `json:"summary"`

	InReplyTo *string `json:"inReplyTo"`

	Published string `json:"published"`

	// this url is for browser to access directly
	URL string `json:"url"`

	// user's activitypub ID / url
	AttributedTo string `json:"attributedTo"`

	Visibility string   `json:"-"`
	To         []string `json:"to" gorm:"-"`
	Cc         []string `json:"cc" gorm:"-"`

	Sensitive bool `json:"sensitive"`

	// same to ID
	AtomUri string `json:"atomUri" gorm:"-"`
	// same to inReplyTo
	InReplyToAtomUri *string `json:"inReplyToAtomUri" gorm:"-"`

	// not implemented
	Conversation string `json:"conversation"`

	// rawContent should be marsharlled later.
	// RawContent string `json:"-"`
	// should be html marshaled
	Content string `json:"content"`

	// not implemented
	ContentMap any `json:"contentMap" gorm:"-"`

	// todo
	// attachments, not implicated
	Attachment []any `json:"attachment" gorm:"-"`

	// if you @someone
	// json because it is hard to parse.
	Tag []*Mention `json:"tag" gorm:"-"`
	// Tag []*Mention `json:"tag" gorm:"serializer:json"`

	Replies *Collection `json:"replies" gorm:"-"`
	// Replies *Collection `json:"replies" gorm:"foreignKey:ID;references:NoteID"`
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

// not viewed
func (o *Note) Autofill() {
	o.Context = NoteContext
	o.Type = TypeNote
	if o.AtomUri == "" {
		o.AtomUri = o.ID
	}
	if o.InReplyToAtomUri == nil {
		o.InReplyToAtomUri = o.InReplyTo
	}
	// for content
	// o.Content = o.RawContent
	name, host, timestamp := utils.ParseStatusesUriToNameHostTimestamp(o.ID)
	if o.AttributedTo == "" {
		o.AttributedTo = utils.NameAndHost2ProfileUrlActivitypubID(name, host)
	}

	if o.URL == "" {
		o.URL = utils.ParseStatusesURL(name, host, timestamp)
	}
	switch o.Visibility {
	case VisiblityPublic:
		o.To = EndpointPublic
		o.Cc = []string{endpointFollower(name, host)}
	case VisiblityUnlisted:
		o.To = []string{endpointFollower(name, host)}
		o.Cc = EndpointPublic
	case VisiblityPrivate:
		o.To = []string{endpointFollower(name, host)}
	case VisiblityDirect:
	default:
		o.Visibility = utils.ParseVisibility(o.To, o.Cc)
	}
	_, _, id := utils.ParseStatusesUriToNameHostTimestamp(o.ID)
	o.Conversation = "tag:fedi.moonchan.xyz," + o.Published[0:10] + ":objectId=" + id + ":objectType=Conversation"
	o.ContentMap = map[string]string{
		"zh": o.Content,
	}
	o.Attachment = []any{}
	o.Tag = []*Mention{}
	o.Replies = &Collection{
		ID:   o.ID + "/replies",
		Type: "Collection",
		First: &CollectionPage{
			Type:   "CollectionPage",
			Next:   o.ID + "/replies" + "?only_other_accounts=true&page=true",
			PartOf: o.ID + "/replies",
			Items:  []any{},
		},
	}

}

func (o *Note) ClearContext() {
	o.Context = nil
}
func (o *Note) GetID() string {
	return o.ID
}
func (o *Note) GetTo() []string {
	return o.To
}
func (o *Note) GetCc() []string {
	return o.Cc
}
func (o *Note) GetActor() string {
	return o.AttributedTo
}

func endpointFollower(username, host string) string {
	return utils.NameAndHost2ProfileUrlActivitypubID(username, host) + "/followers"
}
