package dao

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func CreateStatus(tx *gorm.DB, status *entities.Status) error {

	if err := DB.Create(tx, status); err != nil {
		log.Println(err)
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: status.AttributedTo}, 1); err != nil {
		log.Println(err)
		return err
	}

	tx.Commit()

	return tx.Error
}

// entities.Status
func ReadStatuses(tx *gorm.DB, status *entities.Status) (err error) {

	tx.Preload("Account").Take(status) // it should be the foreign key's var name

	return tx.Error
}

func DeleteStatus(tx *gorm.DB, id string) error {

	status := &entities.Status{
		Id: id,
	}
	if err := DB.Read(tx, status); err != nil {
		log.Println(err)
		return err
	}

	status.DeletedAt = utils.NewTimestamp(true)
	if err := DB.Update(tx, status); err != nil {
		log.Println(err)
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: status.AttributedTo}, -1); err != nil {
		log.Println(err)
		return err
	}

	tx.Commit()

	return tx.Error
}
