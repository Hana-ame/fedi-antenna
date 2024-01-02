package model

import activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"

type User struct {
	// meta
	Email    string `json:"email"`
	PassHash string `json:"passhash"`

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
	Published string `json:"published"`

	AlsoKnownAs []string `json:"alsoKnownAs,omitempty" gorm:"serializer:json"`

	PublicKey *activitypub.PublicKey `json:"publicKey" gorm:"foreignKey:Owner;references:ID"`

	// the url of Image
	IconURL string `json:"-"`
}
