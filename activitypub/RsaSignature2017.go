package activitypub

import (
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/iancoleman/orderedmap"
)

// decryped.
func SignatureObj(user, host string, usTimestamp int64) *orderedmap.OrderedMap {
	o := utils.CreateOrderedMap([]*utils.KV{
		{Key: "type", Value: "RsaSignature2017"},
		{Key: "creator", Value: APUserID(user, host)},
		{Key: "created", Value: utils.TimestampToRFC3339(usTimestamp)},
		{Key: "signatureValue", Value: utils.PrikeyToPubkeyPem(user)},
	})
	return o
}

func PublicKeyObj(user, host string, pkPem string) *orderedmap.OrderedMap {
	o := CreateOrderedMap([]*KV{
		{"id", APUserID(user, host) + "#main-key"},
		{"owner", APUserID(user, host)},
		{"publicKeyPem", pkPem},
	})
	return o
}
