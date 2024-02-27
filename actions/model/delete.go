package model

type Delete struct {
	Context any    `json:"@context,omitempty"`
	ID      string `json:"id"`
	Type    string `json:"type" gorm:"-"`

	Actor string   `json:"actor"`
	To    []string `json:"to" gorm:"serializer:json"`
	// only url
	Object string `json:"object" gorm:"primarykey"`

	SignatureCreator string     `json:"-"`
	Signature        *Signature `json:"omitempty" gorm:"foreignKey:SignatureCreator;references:Creator"`
}

var DeleteContext = "https://www.w3.org/ns/activitystreams"

func (o *Delete) Autofill() {
	o.Context = DeleteContext
	o.Type = "Delete"
}

func (o *Delete) GetType() string {
	return o.Type
}

func (o *Delete) GetRemoteActivitypubID() string {
	return o.Object
}
