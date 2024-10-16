package model

import "github.com/Hana-ame/fedi-antenna/utils"

type Reject struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	Type    string `json:"type" gorm:"-"`
	Actor   string `json:"actor"`

	ObjectID string  `json:"-"`
	Object   *Follow `json:"object" gorm:"foreignKey:ObjectID;references:ID"`

	ID string `json:"id" gorm:"primarykey"`

	CreatedAt int64 `json:"-"`
	DeletedAt int64 `json:"-"`
}

// dunno, from misskey
var RejectContext = []any{
	"https://www.w3.org/ns/activitystreams",
	"https://w3id.org/security/v1",
	utils.NewMapFromKV([]*utils.KV{
		{Key: "manuallyApprovesFollowers", Value: "as:manuallyApprovesFollowers"},
		{Key: "sensitive", Value: "as:sensitive"},
		{Key: "Hashtag", Value: "as:Hashtag"},
		{Key: "quoteUrl", Value: "as:quoteUrl"},
		{Key: "toot", Value: "http://joinmastodon.org/ns#"},
		{Key: "Emoji", Value: "toot:Emoji"},
		{Key: "featured", Value: "toot:featured"},
		{Key: "discoverable", Value: "toot:discoverable"},
		{Key: "schema", Value: "http://schema.org#"},
		{Key: "PropertyValue", Value: "schema:PropertyValue"},
		{Key: "value", Value: "schema:value"},
		{Key: "misskey", Value: "https://misskey-hub.net/ns#"},
		{Key: "_misskey_content", Value: "misskey:_misskey_content"},
		{Key: "_misskey_quote", Value: "misskey:_misskey_quote"},
		{Key: "_misskey_reaction", Value: "misskey:_misskey_reaction"},
		{Key: "_misskey_votes", Value: "misskey:_misskey_votes"},
		{Key: "_misskey_summary", Value: "misskey:_misskey_summary"},
		{Key: "isCat", Value: "misskey:isCat"},
		{Key: "vcard", Value: "http://www.w3.org/2006/vcard/ns#"},
	}),
}

func (o *Reject) Autofill() {
	o.Context = RejectContext
	o.Type = "Reject"
	if o.Actor == "" {
		o.Actor = o.Object.Object
	}
	if o.ObjectID == "" && o.Object != nil {
		o.ObjectID = o.Object.GetID()
	}
	o.Object.ClearContext()
}
