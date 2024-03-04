package dao

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func ReadAccount(tx *gorm.DB, id string) (*entities.Account, error) {
	account := &entities.Account{
		Id: id,
	}
	if err := DB.Read(tx, account); err != nil {
		return account, err
	}
	if account.DeletedAt != 0 {
		return account, fmt.Errorf("Tombstone")
	}
	return account, nil
}

func DeleteAccount(tx *gorm.DB, id string) (*entities.Account, error) {
	account := &entities.Account{
		Id: id,
	}
	if err := DB.Read(tx, account); err != nil {
		return account, err
	}
	account.DeletedAt = utils.NewTimestamp(true)
	if err := DB.Update(tx, account); err != nil {
		return account, err
	}
	return account, nil
}

func UpdateAccountStatusesCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	if err := DB.Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.StatusesCount += delta
	acct.LastStatusAt = utils.ParseStringToPointer(utils.TimestampToRFC3339(utils.NewTimestamp(false)), true)

	if err := DB.Update(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateAccountFollowersCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	if err := DB.Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.FollowersCount += delta

	if err := DB.Update(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateAccountFollowingCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	if err := DB.Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.FollowingCount += delta

	if err := DB.Update(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
