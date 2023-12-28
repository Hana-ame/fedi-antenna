package dao

import (
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
)

func ReadActivitypubUser(id string) (user *activitypub.User, err error) {
	user = &activitypub.User{
		ID: id,
	}

	tx := db.Preload("PublicKey").Preload("Image").Take(user)
	err = tx.Error
	
	return
}
