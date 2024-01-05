package model

import "github.com/Hana-ame/fedi-antenna/core/utils"

type PublicKey struct {
	ID string `json:"id"`

	// the url that return by webfinger
	Owner string `json:"owner" gorm:"primarykey"`

	PublicKeyPem string `json:"publicKeyPem"`
	// PrivateKeyPem string `json:"-"`
}

// should never be used?
func NewPublicKey(id string) *PublicKey {
	pk := utils.GeneratePrivateKey()
	// privateKeyPem := utils.MarshalPrivateKey(pk)
	publicKeyPem := utils.MarshalPublicKey(&pk.PublicKey)
	return &PublicKey{
		ID:           id + "#main-key",
		Owner:        id,
		PublicKeyPem: publicKeyPem,
		// PrivateKeyPem: privateKeyPem,
	}
}
