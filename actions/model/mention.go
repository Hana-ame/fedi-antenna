package model

type Mention struct {
	// "Mention"
	Type string `json:"type"`

	// activitypub ID / url
	Href string `json:"href"`

	// @username@host.domain
	Name string `json:"name" gorm:"primarykey"`
}

func (o *Mention) Autofill() {
	o.Type = TypeMention
}
