package dao

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func UpdateAccountStatusesCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	if err := Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.StatusesCount += delta
	acct.LastStatusAt = utils.ParseStringToPointer(utils.TimestampToRFC3339(utils.NewTimestamp(false)), true)

	if err := Update(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateAccountFollowersCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	if err := Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.FollowersCount += delta

	if err := Update(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateAccountFollowingCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	if err := Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.FollowingCount += delta

	if err := Update(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
