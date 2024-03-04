package dao

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Liked(object, actor string) (liked bool, err error) {

	if err = Read(db, &model.LocalNotify{
		Actor:  actor,
		Object: object,
		Type:   model.NotifyTypeLike,
	}); err != nil {
		log.Printf("%s", err.Error())
		return
	}

	liked = true
	return
}

// use CreateNotify instead
func Favourite(id, object, actor string) error {
	tx := db.Begin()

	status := &entities.Status{Uri: object}
	if err := Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	favourite := &model.LocalNotify{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   model.NotifyTypeLike,

		Owner: status.AttributedTo,
	}

	if err := Create(tx, favourite); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

// use DeleteNotify instead
func Unfavourite(object, actor string) error {
	tx := db.Begin()

	liked, err := Liked(object, actor)
	if err != nil {
		return err
	}
	if !liked {
		return fmt.Errorf("done")
	}

	if err := Delete(tx, &model.LocalNotify{
		Actor:  actor,
		Object: object,
		Type:   model.NotifyTypeLike,
	}); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func Reblogged(object, actor string) (notify *model.LocalNotify, err error) {
	notify = &model.LocalNotify{
		Actor:  actor,
		Object: object,
		Type:   model.NotifyTypeAnnounce,
	}
	err = Read(db, notify)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	return
}

// use CreateNotify instead
func Reblog(id, object, actor, visibility string) error {
	tx := db.Begin()

	status := &entities.Status{Uri: object}
	if err := Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	notify := &model.LocalNotify{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   model.NotifyTypeAnnounce,

		Visibility: visibility,
		Owner:      status.AttributedTo,
	}

	if err := Create(tx, notify); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: actor}, 1); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

// use DeleteNotify instead
func Unreblog(reblogged *model.LocalNotify) error {
	tx := db.Begin()
	reblogged.DeleteAt = utils.NewTimestamp(true)
	if err := Update(tx, reblogged); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: reblogged.Actor}, -1); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func CreateNotify(notify *model.LocalNotify) error {
	tx := db.Begin()
	// notify.DeleteAt = utils.NewTimestamp(true)
	if err := Create(tx, notify); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	if notify.Type == model.NotifyTypeAnnounce {
		if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: notify.Actor}, 1); err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return tx.Error
}

func DeleteNotify(notify *model.LocalNotify) error {
	tx := db.Begin()
	notify.DeleteAt = utils.NewTimestamp(true)
	if err := Update(tx, notify); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	if notify.Type == model.NotifyTypeAnnounce {
		if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: notify.Actor}, -1); err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return tx.Error
}
