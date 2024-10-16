package model

import "github.com/Hana-ame/fedi-antenna/utils"

type Block struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	ID      string `json:"id" gorm:"primarykey"`
	Type    string `json:"type" gorm:"-"`

	// activitypub ID / url
	Actor  string `json:"actor"`
	Object string `json:"object"`

	CreatedAt int64 `json:"-"`
	DeletedAt int64 `json:"-"`
}

var BlockContext = []any{
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

func (o *Block) Autofill() {
	o.Context = BlockContext
	o.Type = "Block"
}

func (o *Block) GetID() string {
	return o.ID
}

func (o *Block) GetType() string {
	return o.Type
}

func (o *Block) GetActor() string {
	return o.Actor
}

func (o *Block) ClearContext() {
	o.Context = nil
}

func (o *Block) GetEndpoint() string {
	return o.Object
}
