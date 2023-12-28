package model

// used in user.tags
type Tags struct {
	Context    string `json:"@context,omitempty" gorm:"-"`
	ID         string `json:"id" gorm:"primarykey"`
	Type       string `json:"type"`
	TotalItems int    `json:"totalItems"`
	Items      []any  `json:"items" gorm:"-"` // todo
}
