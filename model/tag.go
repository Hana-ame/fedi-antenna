package model

type Tag struct {
	// a url
	ID string `json:"id"`

	// emoji
	Type string `json:"type"`

	// :name:
	Name string `json:"name"`

	//RFC
	Updated string `json:"updated"`

	IconURL string `json:"-"`
	Icon    *Image `json:"icon"`
}
