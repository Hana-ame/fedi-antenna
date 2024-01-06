package model

type LocalUser struct {
	// meta
	Email      string `json:"email"`
	PasswdHash string `json:"-"`

	// it is a url
	ID string `json:"id" gorm:"primarykey"`

	// without @host
	PreferredUsername string `json:"preferredUsername" gorm:"type:text collate nocase"`

	// could be empty
	// the name which shows on profile
	Name string `json:"name"`

	Summary string `json:"summary"`

	ManuallyApprovesFollowers bool `json:"manuallyApprovesFollowers"`
	Discoverable              bool `json:"discoverable"`

	// TimestampToRFC3339
	Published int64 `json:"published"`

	AlsoKnownAs []string `json:"alsoKnownAs,omitempty" gorm:"serializer:json"`

	PublicKeyPem  string
	PrivateKeyPem string

	// the url of Image, avatar
	IconURL string `json:"-"`
	// the url of Image, background
	ImageURL string `json:"-"`

	DeletedAt int64
}
