package dao

import (
	"crypto/rsa"
	"log"

	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

// activitypub.User
func ReadActivitypubUser(id string) (user *activitypub.User, err error) {
	user = &activitypub.User{
		ID: id,
	}

	tx := db.Preload("Icon").Preload("PublicKey").Take(user) // it should be the foreign key's var name
	err = tx.Error

	user.Autofill()

	return
}

// what is it???
// find in local
// if not found then fetch from remote
func ReadPublicKeyByOwner(id string) (pk *activitypub.PublicKey, err error) {
	tx := db.Begin()

	pk = &activitypub.PublicKey{
		Owner: id,
	}
	err = Read(tx, pk)
	if err == nil {
		return
	}

	return
}

// the most direct way to read local user's privateKey
func ReadPrivateKeyByOwner(id string) (pk *rsa.PrivateKey, err error) {
	tx := db.Begin()

	lu := &core.LocalUser{
		ActivitypubID: id,
	}
	err = Read(tx, lu)
	if err != nil {
		log.Println(err)
		return
	}
	pk, err = utils.ParsePrivateKey(lu.PrivateKeyPem)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func UpdateAccountStatusesCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	// tx := db.Begin()

	if err := Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.StatusesCount += delta
	acct.LastStatusAt = utils.ParseStringToPointer(utils.TimestampToRFC3339(utils.Now()), true)

	if err := Update(tx, acct); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}

func UpdateAccountFollowersCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	// tx := db.Begin()

	if err := Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.FollowersCount += delta

	if err := Update(tx, acct); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}

func UpdateAccountFollowingCount(tx *gorm.DB, acct *entities.Account, delta int) error {
	// tx := db.Begin()

	if err := Read(tx, acct); err != nil {
		log.Println(err)
		return err
	}

	acct.FollowingCount += delta

	if err := Update(tx, acct); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}
