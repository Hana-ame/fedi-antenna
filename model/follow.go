package model

import "github.com/Hana-ame/fedi-antenna/core/utils"

type Follow struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	ID      string `json:"id" gorm:"primarykey"`
	Type    string `json:"type" gorm:"-"`

	// activitypub ID / url
	Actor  string `json:"actor"`
	Object string `json:"object"`
}

var FollowContext = []any{
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

func (o *Follow) Autofill() {
	o.Context = FollowContext
	o.Type = "Follow"
}

func (o *Follow) GetID() string {
	return o.ID
}

func (o *Follow) GetType() string {
	return o.Type
}

func (o *Follow) GetActor() string {
	return o.Actor
}

func (o *Follow) GetObject() string {
	return o.Object
}

func (o *Follow) ClearContext() {
	o.Context = nil
}
