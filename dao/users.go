package dao

import (
	"crypto/rsa"
	"log"

	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// activitypub.User
func ReadActivitypubUser(id string) (user *activitypub.User, err error) {
	user = &activitypub.User{
		ID: id,
	}

	tx := db.Preload("Icon").Preload("PublicKey").Take(user) // it should be the foreign key's var name
	err = tx.Error

	user.Autofill()

	return
}

// what is it???
// find in local
// if not found then fetch from remote
func ReadPublicKeyByOwner(id string) (pk *activitypub.PublicKey, err error) {
	pk = &activitypub.PublicKey{
		Owner: id,
	}
	err = Read(pk)
	if err == nil {
		return
	}

	return
}

// the most direct way to read local user's privateKey
func ReadPrivateKeyByOwner(id string) (pk *rsa.PrivateKey, err error) {
	lu := &core.LocalUser{
		ID: id,
	}
	err = Read(lu)
	if err != nil {
		log.Println(err)
		return
	}
	pk, err = utils.ParsePrivateKey(lu.PrivateKeyPem)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
