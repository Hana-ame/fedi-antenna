package model

import "github.com/Hana-ame/fedi-antenna/core/utils"

// no need to save to db
type Undo struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	Type    string `json:"type" gorm:"-"`
	ID      string `json:"id" gorm:"primarykey"`
	Actor   string `json:"actor"`

	ObjectID string   `json:"-"`
	Object   Sendable `json:"object" gorm:"-"`

	// RFC
	Published string `json:"pbulished"`
}

var UndoContext = []any{
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

func (o *Undo) Autofill() {
	o.Context = UndoContext
	o.Type = "Undo"
	o.ID = o.Object.GetID() + "/undo"
	o.Actor = o.Object.GetActor()
	o.ObjectID = o.Object.GetID()
	if o.Published == "" {
		o.Published = utils.MicroSecondToRFC3339(utils.Now())
	}
	o.Object.ClearContext()
}

func (o *Undo) GetType() string {
	return o.Type
}

func (o *Undo) GetActor() string {
	return o.Object.GetActor()
}
func (o *Undo) GetObject() string {
	return o.Object.GetObject()
}
