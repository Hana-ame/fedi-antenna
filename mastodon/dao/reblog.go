package dao

import (
	"log"
	"strconv"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func CreateReblog(tx *gorm.DB, object, actor, visibility string) error {
	status := &entities.Status{Uri: object}
	if err := DB.Read(tx, status); err != nil {
		log.Println(err)
		return err
	}

	reblog := &Reblog{
		Id:         strconv.Itoa(int(utils.NewTimestamp(false))),
		Actor:      actor,
		Object:     object,
		Owner:      status.AttributedTo,
		Visibility: visibility,
		DeletedAt:  -1,
	}
	if err := DB.Create(tx, reblog); err != nil {
		log.Println(err)
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: reblog.Actor}, 1); err != nil {
		log.Println(err)
		return err
	}

	tx.Commit()

	return tx.Error
}

func DeleteReblog(tx *gorm.DB, object, actor string) error {

	reblog := &Reblog{
		Actor:     actor,
		Object:    object,
		DeletedAt: -1,
	}
	if err := DB.Read(tx, reblog); err != nil {
		log.Println(err)
		return err
	}
	reblog.DeletedAt = utils.NewTimestamp(true)
	if err := DB.Update(tx, reblog); err != nil {
		log.Println(err)
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: actor}, -1); err != nil {
		log.Println(err)
		return err
	}

	tx.Commit()

	return tx.Error
}
