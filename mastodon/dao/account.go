package dao

import (
	"fmt"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func ReadAccount(tx *gorm.DB, acct *entities.Account) error {

	if err := Read(tx, acct); err != nil {
		return err
	}

	if acct.DeletedAt != 0 {
		return fmt.Errorf("Tombstone")
	}

	return nil
}

func DeleteAccount(tx *gorm.DB, acct *entities.Account) (*entities.Account, error) {

	acct.DeletedAt = utils.Timestamp(true)

	if err := Update(tx, acct); err != nil {
		return acct, err
	}

	return acct, nil
}

func UpdateAccountStatusesCount(tx *gorm.DB, acct *entities.Account, delta int) error {

	if err := Read(tx, acct); err != nil {
		logE(err)
		return err
	}

	acct.StatusesCount += delta
	if delta > 0 {
		acct.LastStatusAt = utils.ParseStringToPointer(utils.TimestampToRFC3339(utils.Timestamp(false)), true)
	}

	if err := Update(tx, acct); err != nil {
		logE(err)
		return err
	}

	return nil
}

func UpdateAccountFollowersCount(tx *gorm.DB, acct *entities.Account, delta int) error {

	if err := Read(tx, acct); err != nil {
		logE(err)
		return err
	}

	acct.FollowersCount += delta

	if err := Update(tx, acct); err != nil {
		logE(err)
		return err
	}

	return nil
}

func UpdateAccountFollowingCount(tx *gorm.DB, acct *entities.Account, delta int) error {

	if err := Read(tx, acct); err != nil {
		logE(err)
		return err
	}

	acct.FollowingCount += delta

	if err := Update(tx, acct); err != nil {
		logE(err)
		return err
	}

	return nil
}
