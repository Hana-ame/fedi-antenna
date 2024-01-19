package dao

import (
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// entities.Status
func ReadMastodonStatuses(status *entities.Status) (err error) {

	tx := db.Preload("Account").Take(status) // it should be the foreign key's var name
	err = tx.Error

	return
}
