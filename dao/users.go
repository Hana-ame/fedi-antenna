package dao

import (
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
)

func ReadActivitypubUser(id string) (user *activitypub.User, err error) {
	user = &activitypub.User{
		ID: id,
	}

	err = Read(user)
	if err == nil {
		user.Autofill()
		user.PublicKey = &activitypub.PublicKey{
			Owner: user.ID,
		}
		Read(user.PublicKey)
		user.Icon = &activitypub.Image{
			URL: user.IconURL,
		}
		Read(user.Icon)
	}

	return
}
