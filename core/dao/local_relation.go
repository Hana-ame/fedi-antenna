package dao

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
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
		_, host := utils.ActivitypubID2NameAndHost(r.Actor)
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

// actor's following, object's follower
func UpdateAccountFollowerFollowingCount(tx *gorm.DB, object, actor string, delta int) (err error) {
	err = UpdateAccountFollowingCount(tx, &entities.Account{Uri: actor}, delta)
	if err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}
	err = UpdateAccountFollowersCount(tx, &entities.Account{Uri: object}, delta)
	if err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}
	return nil
}

// set to padding
func Follow(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := ReadRelationship(object, actor)
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

	relationship, err := ReadRelationship(object, actor)
	if err != nil {
		return err
	}
	if !(relationship.Following || relationship.Requested) {
		return fmt.Errorf("done")
	} else if relationship.BlockedBy || relationship.Blocking {
		return fmt.Errorf("forbidden")
	}

	lr := &model.LocalRelation{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   model.RelationTypeNone,
		Status: model.RelationStatusUndo,
	}
	if err := Update(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}
	if relationship.Following {
		err = UpdateAccountFollowerFollowingCount(tx, actor, object, -1)
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

	relationship, err := ReadRelationship(object, actor)
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
		err = UpdateAccountFollowerFollowingCount(tx, actor, object, 1)
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

	relationship, err := ReadRelationship(object, actor)
	if err != nil {
		return err
	}
	if !(relationship.RequestedBy || relationship.FollowedBy) {
		return fmt.Errorf("forbidden")
	}

	lr := &model.LocalRelation{
		ID:     id,
		Actor:  object,
		Object: actor,
		Type:   model.RelationTypeNone,
		Status: model.RelationStatusRejected,
	}
	if err := Update(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	if relationship.FollowedBy {
		err = UpdateAccountFollowerFollowingCount(tx, actor, object, -1)
		if err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return tx.Error
}

func Block(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := ReadRelationship(object, actor)
	if err != nil {
		return err
	}
	if relationship.Blocking {
		return fmt.Errorf("done")
	}

	if relationship.FollowedBy {
		err = UpdateAccountFollowerFollowingCount(tx, actor, object, -1)
		if err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
	}
	if relationship.Following {
		err = UpdateAccountFollowerFollowingCount(tx, object, actor, -1)
		if err != nil {
			log.Printf("%s", err.Error())
			tx.Rollback()
			return err
		}
	}
	// err := Delete(tx, &model)
	lr := &model.LocalRelation{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   model.RelationTypeBlock,
		Status: model.RelationStatusBlocking,
	}

	if err := Update(tx, lr); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func Unblock(id, object, actor string) error {
	tx := db.Begin()

	relationship, err := ReadRelationship(object, actor)
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
		Status: model.RelationStatusUnblocked,
	}

	if err := Update(tx, lr); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}
