package model

type Collection struct {
	// now it only occurs at note. so that is ok?
	NoteID string          `json:"-"`
	ID     string          `json:"id" gorm:"primarykey"`
	Type   string          `json:"type"`
	First  *CollectionPage `json:"first" gorm:"serializer:json"`
}

type CollectionPage struct {
	Type   string `json:"type"`
	Next   string `json:"next"`
	PartOf string `json:"partOf"`

	// unknown.
	Items []any `json:"items" gorm:"serializer:json"`
}
