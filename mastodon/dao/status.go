package dao

import (
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func CreateStatus(tx *gorm.DB, status *entities.Status) error {

	if err := Create(tx, status); err != nil {
		logE(err)
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: status.AttributedTo}, 1); err != nil {
		logE(err)
		return err
	}

	return tx.Error
}

// entities.Status
func ReadStatuses(tx *gorm.DB, status *entities.Status) (err error) {

	tx.Preload("Account").Preload("Reblog").Take(status) // it should be the foreign key's var name

	return tx.Error
}

func DeleteStatus(tx *gorm.DB, status *entities.Status) error {

	status.DeletedAt = utils.Timestamp(true)
	if err := Update(tx, status); err != nil {
		logE(err)
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: status.AttributedTo}, -1); err != nil {
		logE(err)
		return err
	}

	return tx.Error
}

func ReadPublicStatus(tx *gorm.DB, limit int) (statuses []*entities.Status, err error) {

	if limit > 40 {
		limit = 40
	} else if limit <= 0 {
		limit = 20
	}

	// not tested
	tx.Raw("SELECT * FROM statuses WHERE visibility = 'public' AND reblog_uri IS NULL ORDER BY create_at DESC LIMIT ?", limit).Scan(&statuses)
	err = tx.Error

	return
}
