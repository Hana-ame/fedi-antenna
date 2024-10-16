package model

type Accept struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	Type    string `json:"type" gorm:"-"`
	Actor   string `json:"actor"`

	ObjectID string  `json:"-"`
	Object   *Follow `json:"object" gorm:"foreignKey:ObjectID;references:ID"`

	ID string `json:"id" gorm:"primarykey"`

	CreatedAt int64 `json:"-"`
	DeletedAt int64 `json:"-"`
}

// dunno, from mastodon
var AcceptContext = "https://www.w3.org/ns/activitystreams"

func (o *Accept) Autofill() {
	o.Context = AcceptContext
	o.Type = "Accept"
	if o.Actor == "" {
		o.Actor = o.Object.Object
	}
	if o.ObjectID == "" && o.Object != nil {
		o.ObjectID = o.Object.GetID()
	}
	o.Object.ClearContext()
}
