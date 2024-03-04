package dao

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func ReadRelationship(tx *gorm.DB, object, actor string) (*entities.Relationship, error) {

	account := &entities.Account{Uri: object}
	if err := DB.Read(tx, account); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	a2o := &Relation{
		Actor:  actor,
		Object: object,
	}
	if err := DB.Read(tx, a2o); err != nil {
		// log.Printf("%s", err.Error())
		// return nil, err
	}
	// object to actor
	o2a := &Relation{
		Actor:  object,
		Object: actor,
	}
	if err := DB.Read(tx, o2a); err != nil {
		// log.Printf("%s", err.Error())
		// return nil, err
	}

	relationship := &entities.Relationship{
		// Type: String (cast from an integer, but not guaranteed to be a number)
		// Description: The account ID.
		Attributes: (account.Id),
		// Type: Boolean
		// Description: Are you following this user?
		Following: a2o.Status == RelationStatusAccepted,
		// Type: Boolean
		// Description: Are you receiving this user’s boosts in your home timeline?
		ShowingReblogs: true,
		// Type: Boolean
		// Description: Have you enabled notifications for this user?
		Notifying: false,
		// Type: Array of String (ISO 639-1 language two-letter code)
		// Description: Which languages are you following from this user?
		Languages: []string{},
		// Type: Boolean
		// Description: Are you followed by this user?
		FollowedBy: o2a.Status == RelationStatusAccepted,
		// Type: Boolean
		// Description: Are you blocking this user?
		Blocking: a2o.Status == RelationStatusBlocking,
		// Type: Boolean
		// Description: Is this user blocking you?
		BlockedBy: o2a.Status == RelationStatusBlocking,
		// Type: Boolean
		// Description: Are you muting this user?
		Muting: false,
		// Type: Boolean
		// Description: Are you muting notifications from this user?
		MutingNotifications: false,
		// Type: Boolean
		// Description: Do you have a pending follow request for this user?
		Requested: a2o.Status == model.RelationStatusPadding,
		// Type: Boolean
		// Description: Has this user requested to follow you?
		RequestedBy: o2a.Status == model.RelationStatusPadding,
		// Type: Boolean
		// Description: Are you blocking this user’s domain?
		DomainBlocking: false,
		// Type: Boolean
		// Description: Are you featuring this user on your profile?
		Endorsed: false,
		// Type: String
		// Description: This user’s profile bio
		Note: account.Note,
	}
	return relationship, nil
}

// actor's following, object's follower
func UpdateAccountFollowerFollowingCount(tx *gorm.DB, object, actor string, delta int) (err error) {
	err = UpdateAccountFollowingCount(tx, &entities.Account{Uri: actor}, delta)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	err = UpdateAccountFollowersCount(tx, &entities.Account{Uri: object}, delta)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	return nil
}

func Follow(tx *gorm.DB, object, actor string) error {
	relationship, err := ReadRelationship(tx, object, actor)
	if err != nil {
		return err
	} else if relationship.Following || relationship.Requested {
		return fmt.Errorf("done")
	} else if relationship.BlockedBy || relationship.Blocking {
		return fmt.Errorf("forbidden")
	}

	lr := &Relation{
		Actor:  actor,
		Object: object,
		Status: model.RelationStatusPadding,
	}

	if err := DB.Create(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return err
	}

	tx.Commit()

	return tx.Error
}

func Unfollow(tx *gorm.DB, object, actor string) error {
	relationship, err := ReadRelationship(tx, object, actor)
	if err != nil {
		return err
	} else if !(relationship.Following || relationship.Requested) {
		return fmt.Errorf("done")
	} else if relationship.BlockedBy || relationship.Blocking {
		return fmt.Errorf("forbidden")
	}

	lr := &Relation{
		Actor:  actor,
		Object: object,
	}
	if err := DB.Delete(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return err
	}
	if relationship.Following {
		err = UpdateAccountFollowerFollowingCount(tx, actor, object, -1)
		if err != nil {
			log.Printf("%s", err.Error())
			return err
		}
	}

	tx.Commit()

	return tx.Error
}

// actor accpet object's request
func Accept(tx *gorm.DB, object, actor string) error {

	relationship, err := ReadRelationship(tx, object, actor)
	if err != nil {
		return err
	} else if !relationship.RequestedBy {
		return fmt.Errorf("forbidden")
	} else if relationship.BlockedBy || relationship.Blocking {
		return fmt.Errorf("forbidden")
	}

	lr := &Relation{
		Actor:  object,
		Object: actor,
		Status: model.RelationStatusAccepted,
	}
	if err := DB.Update(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return err
	}
	if relationship.RequestedBy {
		err = UpdateAccountFollowerFollowingCount(tx, actor, object, 1)
		if err != nil {
			log.Printf("%s", err.Error())
			return err
		}
	}

	tx.Commit()

	return tx.Error
}

// actor reject object's request
func Reject(tx *gorm.DB, object, actor string) error {

	relationship, err := ReadRelationship(tx, object, actor)
	if err != nil {
		return err
	} else if !(relationship.RequestedBy || relationship.FollowedBy) {
		return fmt.Errorf("forbidden")
	} else if relationship.BlockedBy || relationship.Blocking {
		return fmt.Errorf("forbidden")
	}

	lr := &Relation{
		Actor:  object,
		Object: actor,
	}
	if err := DB.Delete(tx, lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return err
	}

	if relationship.FollowedBy {
		err = UpdateAccountFollowerFollowingCount(tx, actor, object, -1)
		if err != nil {
			log.Printf("%s", err.Error())
			return err
		}
	}

	tx.Commit()

	return tx.Error
}

func Block(tx *gorm.DB, object, actor string) error {

	relationship, err := ReadRelationship(tx, object, actor)
	if err != nil {
		return err
	} else if relationship.Blocking {
		return fmt.Errorf("done")
	}

	if relationship.FollowedBy || relationship.RequestedBy {
		lr := &Relation{
			Actor:  object,
			Object: actor,
		}
		if err := DB.Delete(tx, lr); err != nil {
			// should not
			log.Printf("%s", err.Error())
			return err
		}
	}

	if relationship.FollowedBy {
		if err := UpdateAccountFollowerFollowingCount(tx, actor, object, -1); err != nil {
			log.Printf("%s", err.Error())
			return err
		}
	}
	if relationship.Following {
		if err := UpdateAccountFollowerFollowingCount(tx, object, actor, -1); err != nil {
			log.Printf("%s", err.Error())
			return err
		}
	}
	// err := Delete(tx, &model)
	lr := &Relation{
		Actor:  actor,
		Object: object,
		Status: model.RelationStatusBlocking,
	}

	if err := DB.Update(tx, lr); err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	tx.Commit()

	return tx.Error
}

func Unblock(tx *gorm.DB, object, actor string) error {

	relationship, err := ReadRelationship(tx, object, actor)
	if err != nil {
		return err
	} else if !relationship.Blocking {
		return fmt.Errorf("done")
	}

	lr := &Relation{
		Actor:  actor,
		Object: object,
		Status: RelationStatusBlocking,
	}
	if err := DB.Delete(tx, lr); err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	tx.Commit()

	return tx.Error
}
