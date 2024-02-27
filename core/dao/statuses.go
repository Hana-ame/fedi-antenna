package dao

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// entities.Status
func ReadMastodonStatuses(status *entities.Status) (err error) {

	tx := db.Preload("Account").Take(status) // it should be the foreign key's var name
	err = tx.Error

	return
}

func DeleteStatus(status *entities.Status) (err error) {
	if err = Read(status); err != nil {
		log.Println(err)
		return
	}
	if err = Delete(status); err != nil {
		log.Println(err)
		return
	}
	UpdateAccountStatusesCount(&entities.Account{Uri: status.AttributedTo}, -1)

	return
}

func CreateStatus(status *entities.Status) (err error) {
	if err = Create(status); err != nil {
		log.Println(err)
		return
	}
	UpdateAccountStatusesCount(&entities.Account{Uri: status.AttributedTo}, 1)
	return
}
