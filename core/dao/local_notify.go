package dao

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/core/model"
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

func Unfavourite(id, object, actor string) error {
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

func Reblogged(object, actor string) (reblogged bool, err error) {

	if err = Read(db, &model.LocalNotify{
		Actor:  actor,
		Object: object,
		Type:   model.NotifyTypeAnnounce,
	}); err != nil {
		log.Printf("%s", err.Error())
		return
	}

	reblogged = true
	return
}

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

func Unreblog(object, actor string) error {
	tx := db.Begin()

	reblogged, err := Reblogged(object, actor)
	if err != nil {
		return err
	}
	if !reblogged {
		return fmt.Errorf("done")
	}

	if err := Delete(tx, &model.LocalNotify{
		Actor:  actor,
		Object: object,
		Type:   model.NotifyTypeAnnounce,
	}); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	if err := UpdateAccountStatusesCount(tx, &entities.Account{Uri: actor}, -1); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}
