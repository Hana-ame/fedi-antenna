package model

// used in user.icon
type Emoji struct {
	ID   string `json:"id" gorm:"primarykey"`
	Type string `json:"type"`
	Name string `json:"name"`

	// TimestampToRFC3339
	Updated string `json:"updated"`

	IconURL string `json:"-"`
	Icon    *Image `json:"icon" gorm:"foreignKey:IconURL;references:URL"`
}
