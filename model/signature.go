package model

// no id?

type Signature struct {
	Type           string `json:"type"`
	Creator        string `json:"creator"`
	Created        string `json:"created"`
	SignatureValue string `json:"signatureValue"`
}
