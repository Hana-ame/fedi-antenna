package dao

import (
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// activitypub.User
func ReadActivitypubUser(user *activitypub.User) error {

	tx := db.Preload("Icon").Preload("PublicKey").Take(user) // it should be the foreign key's var name

	return tx.Error
}

func DeletePerson(id string) error {
	Delete(db, &entities.Account{Uri: id})
	Delete(db, &entities.Status{AttributedTo: id})
	return nil
}
