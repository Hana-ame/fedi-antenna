package model

type Note struct {
	// not used?
	Context any `json:"omitempty" gorm:"-"`

	//
	ID string `json:"id" gorm:"primarykey"`

	// alwasy "Note"
	Type string `json:"type" gorm:"-"`

	// user's activitypub ID / url
	AttributeTo string `json:"attributeTo"`

	// this should be html marshaled
	Content string `json:"content"`

	Published string `json:"published"`

	To []string `json:"to" gorm:"serializer:json"`
	Cc []string `json:"cc" gorm:"serializer:json"`

	InReplyTo  *string `json:"inReplyTo"`
	Attachment []any   `json:"attachment" gorm:"-"`
	Sensitive  bool    `json:"sensitive"`
	Tag        []*Tag  `json:"tag" gorm:"serializer:json"`
}

func (o *Note) Autofill() {
	o.ID = "Note"

}
