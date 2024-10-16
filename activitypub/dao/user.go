package dao

import "github.com/Hana-ame/fedi-antenna/activitypub/dao/model"

// activitypub.User
func ReadUser(id string) (user *model.User, err error) {
	user = &model.User{
		ID: id,
	}

	tx := db.Preload("Icon").Preload("PublicKey").Take(user) // it should be the foreign key's var name
	err = tx.Error

	user.Autofill()

	return
}

// create new user.
// func CreateUser(acct, host string) (user *model.User, err error) {
// 	user = model.NewUser(acct, host)

// 	err = Create(user)

// 	return
// }
