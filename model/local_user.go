package model

type LocalUser struct {
	// meta
	Email      string `json:"email" gorm:"index:,unique;type:text collate nocase"`
	PasswdHash string `json:"-"`

	// without @host
	Username string `json:"preferredUsername" gorm:"primarykey;index:,unique;type:text collate nocase"`
	Host     string `gorm:"primarykey;index:,unique;type:text collate nocase"`

	// activitypub url
	ActivitypubID string `gorm:"index:,unique;type:text collate nocase"`
	// timestamp string
	AccountID string `gorm:"index:,unique;type:text collate nocase"`

	AlsoKnownAs []string `json:"alsoKnownAs,omitempty" gorm:"serializer:json"`

	PrivateKeyPem string

	CreatedAt int64
	DeletedAt int64
}
