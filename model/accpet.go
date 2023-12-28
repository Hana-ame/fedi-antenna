package model

type Accept struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	Type    string `json:"type" gorm:"-"`
	Actor   string `json:"actor"`

	ObjectID string   `json:"-"`
	Object   Sendable `json:"object" gorm:"foreignKey:ObjectID;references:ID"`

	ID string `json:"id" gorm:"primarykey"`
}

var AcceptContext = "https://www.w3.org/ns/activitystreams"

func (o *Accept) Autofill() {
	o.Context = AcceptContext
	o.Type = "Accpet"
	o.Actor = o.Object.GetObject()
	o.ObjectID = o.Object.GetID()
	o.Object.ClearContext()
}
