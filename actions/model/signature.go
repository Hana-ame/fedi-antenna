package model

// no id?

// only created at delete user?
type Signature struct {
	Type string `json:"type"`

	// which contains '#main-key'
	Creator string `json:"creator" gorm:"primarykey"`

	//
	Created        string `json:"created"`
	SignatureValue string `json:"signatureValue"`
}
