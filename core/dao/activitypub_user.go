package dao

import (
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
)

// activitypub.User
func ReadActivitypubUser(user *activitypub.User) error {

	tx := db.Preload("Icon").Preload("PublicKey").Take(user) // it should be the foreign key's var name

	return tx.Error
}

// please do not use it.
// use read instead.

// what is it???
// find in local
// if not found then fetch from remote
// func ReadPublicKeyByOwner(id string) (pk *activitypub.PublicKey, err error) {
// 	// tx := db.Begin()

// 	pk = &activitypub.PublicKey{
// 		Owner: id,
// 	}
// 	err = Read(db, pk)

// 	return
// }
