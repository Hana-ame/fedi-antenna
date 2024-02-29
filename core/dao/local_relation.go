package dao

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// 这是啥
// 不想改表了，很麻烦。
// not tested.
func ReadFollowersByLocaluserID(id string) (sharedInboxes []string, err error) {
	// 1. query
	var relations []model.LocalRelation
	tx := db.Where("object = ?", id).Find(&relations)
	err = tx.Error
	if err != nil {
		return
	}

	// 2. select hosts.

	// 2.1. remove more than once
	var m map[string]any = make(map[string]any)
	for _, r := range relations {
		_, host := utils.ParseNameAndHost(r.Actor)
		m[host] = true
	}
	// 2.2. host to shared inbox
	var sharedInbox []model.FediStatus
	var hosts []string = make([]string, len(m))
	i := 0
	for k, _ := range m {
		hosts[i] = k
		i++
	}
	db.Find(&sharedInbox, hosts)

	// 3. move into []string
	sharedInboxes = make([]string, len(sharedInbox))
	for i, sharedinbox := range sharedInbox {
		sharedInboxes[i] = sharedinbox.SharedInbox
	}
	return
}

func Relationship(object, actor string) (relationship *entities.Relationship, err error) {

	o2a := &model.LocalRelation{
		Actor:  actor,
		Object: object,
	}
	a2o := &model.LocalRelation{
		Actor:  object,
		Object: actor,
	}
	if err = Read(db, o2a); err != nil {
		return
	}
	if err = Read(db, a2o); err != nil {
		return
	}

	relationship = &entities.Relationship{
		Following:  a2o.Status == model.RelationStatusAccepted,
		FollowedBy: o2a.Status == model.RelationStatusAccepted,

		Requested:   a2o.Status == model.RelationStatusPadding,
		RequestedBy: o2a.Status == model.RelationStatusPadding,

		Blocking:  a2o.Status == model.RelationStatusBlocking,
		BlockedBy: o2a.Status == model.RelationStatusBlocking,
	}

	return
}

// set to padding
func Follow(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := Relationship(object, actor)
	if err != nil {
		return err
	}
	if relationship.Following || relationship.Requested {
		return fmt.Errorf("done")
	}
	if relationship.BlockedBy || relationship.Blocking {
		return fmt.Errorf("forbidden")
	}

	lr := &model.LocalRelation{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   model.RelationTypeFollow,
		Status: model.RelationStatusPadding,
	}

	if err := Create(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func Unfollow(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := Relationship(object, actor)
	if err != nil {
		return err
	}
	if !relationship.Following && !relationship.Requested {
		return fmt.Errorf("done")
	} else if relationship.BlockedBy || relationship.Blocking {
		return fmt.Errorf("forbidden")
	}

	lr := &model.LocalRelation{
		Actor:  actor,
		Object: object,
		Type:   model.RelationTypeFollow,
	}
	if err := Delete(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}
	if relationship.Following {
		err = UpdateAccountFollowingCount(tx, &entities.Account{Uri: actor}, -1)
		if err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
		err = UpdateAccountFollowersCount(tx, &entities.Account{Uri: object}, -1)
		if err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return tx.Error
}

// actor accpet object's request
func Accept(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := Relationship(object, actor)
	if err != nil {
		return err
	}
	if !relationship.RequestedBy {
		return fmt.Errorf("forbidden")
	}

	lr := &model.LocalRelation{
		ID:     id,
		Actor:  object,
		Object: actor,
		Type:   model.RelationTypeFollow,
		Status: model.RelationStatusAccepted,
	}
	if err := Update(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}
	if relationship.RequestedBy {
		err = UpdateAccountFollowingCount(tx, &entities.Account{Uri: lr.Actor}, 1)
		if err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
		err = UpdateAccountFollowersCount(tx, &entities.Account{Uri: lr.Object}, 1)
		if err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return tx.Error
}

// actor reject object's request
func Reject(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := Relationship(object, actor)
	if err != nil {
		return err
	}
	if !relationship.RequestedBy && !relationship.FollowedBy {
		return fmt.Errorf("forbidden")
	}

	lr := &model.LocalRelation{
		ID:     id,
		Actor:  object,
		Object: actor,
		Type:   model.RelationTypeFollow,
	}
	if err := Delete(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func Block(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := Relationship(object, actor)
	if err != nil {
		return err
	}
	if relationship.Blocking {
		return fmt.Errorf("done")
	}

	if relationship.FollowedBy {
		err = Reject("", object, actor)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if relationship.Following {
		err = Unfollow("", object, actor)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	lr := &model.LocalRelation{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   model.RelationTypeBlock,
		Status: model.RelationStatusBlocking,
	}

	if err := Create(tx, lr); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func Unblock(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := Relationship(object, actor)
	if err != nil {
		return err
	}
	if !relationship.Blocking {
		return fmt.Errorf("done")
	}

	lr := &model.LocalRelation{
		Actor:  actor,
		Object: object,
		Type:   model.RelationTypeBlock,
	}

	if err := Delete(tx, lr); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}
