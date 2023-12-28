package model

type Delete struct {
	Context any      `json:"@context"`
	ID      string   `json:"id"`
	Type    string   `json:"type" gorm:"-"`
	Actor   string   `json:"actor"`
	To      []string `json:"to" gorm:"serializer:json"`
	// only url
	Object string `json:"object" gorm:"primarykey"`
}

var DeleteContext = "https://www.w3.org/ns/activitystreams"

func (o *Delete) Autofill() {
	o.Context = DeleteContext
	o.Type = "Delete"
}
