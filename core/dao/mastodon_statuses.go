package dao

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// entities.Status
func ReadMastodonStatuses(status *entities.Status) (err error) {

	tx := db.Preload("Account").Take(status) // it should be the foreign key's var name
	err = tx.Error

	return
}

func DeleteStatus(status *entities.Status) error {
	tx := db.Begin()

	// if err := Read(tx, status); err != nil {
	// 	log.Println(err)
	// 	tx.Rollback()
	// 	return err
	// }
	status.DeletedAt = utils.NewTimestamp(true)
	if err := Update(tx, status); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: status.AttributedTo}, -1); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func CreateStatus(status *entities.Status) error {
	tx := db.Begin()

	if err := Create(tx, status); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: status.AttributedTo}, 1); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}
