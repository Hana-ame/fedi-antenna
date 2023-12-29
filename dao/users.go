package dao

import (
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
)

func ReadActivitypubUser(id string) (user *activitypub.User, err error) {
	user = &activitypub.User{
		ID: id,
	}

	tx := db.Preload("Icon").Preload("PublicKey").Take(user) // it should be the foreign key's var name
	err = tx.Error

	return
}
