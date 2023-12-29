package model

type Accept struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	Type    string `json:"type" gorm:"-"`
	Actor   string `json:"actor"`

	ObjectID string   `json:"-"`
	Object   Sendable `json:"object" gorm:"-"`

	ID string `json:"id" gorm:"primarykey"`
}

// dunno, from mastodon
var AcceptContext = "https://www.w3.org/ns/activitystreams"

func (o *Accept) Autofill() {
	o.Context = AcceptContext
	o.Type = "Accept"
	o.Actor = o.Object.GetObject()
	o.ObjectID = o.Object.GetID()
	o.Object.ClearContext()
}
