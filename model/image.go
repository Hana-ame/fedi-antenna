package model

// used in user.icon
type Image struct {
	Type      string `json:"type"`
	MediaType string `json:"mediaType"`
	URL       string `json:"url" gorm:"primarykey"`
}
