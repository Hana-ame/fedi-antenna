package core

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Block(id, object, actor string) error {
	{
		// delete follow id -> actor
		lr := &core.LocalRelation{
			// ID:     utils.GenerateObjectID(typ, host),
			Actor:  actor,
			Object: object,
			Type:   core.RelationTypeFollow,
			// Status: core.RelationStatusPadding,
		}
		if err := dao.Where(*lr).Find(lr); err == nil {
			if err := dao.Delete(lr); err != nil {
				if lr.Status == core.RelationStatusAccepted {
					dao.UpdateAccountFollowingCount(&entities.Account{Uri: lr.Actor}, -1)
					dao.UpdateAccountFollowersCount(&entities.Account{Uri: lr.Object}, -1)
				}
			} else {
				return err
			}
		}
	}
	{
		// delete follow actor -> id
		lr := &core.LocalRelation{
			// ID:     utils.GenerateObjectID(typ, host),
			Actor:  object,
			Object: actor,
			Type:   core.RelationTypeFollow,
			// Status: core.RelationStatusAccepted,
		}
		if err := dao.Where(*lr).Find(lr); err == nil {
			if err := dao.Delete(lr); err != nil {
				if lr.Status == core.RelationStatusAccepted {
					dao.UpdateAccountFollowingCount(&entities.Account{Uri: lr.Actor}, -1)
					dao.UpdateAccountFollowersCount(&entities.Account{Uri: lr.Object}, -1)
				}
			} else {
				return err
			}
		}
	}
	lr := &core.LocalRelation{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   core.RelationTypeBlock,
		Status: core.RelationStatusBlocked,
	}

	// if err := dao.Where("ID = ?", lr.ID).First(lr); err == nil {
	// 	return fmt.Errorf("already exists")
	// }

	if err := dao.Create(lr); err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	return nil
}

func Unblock(id, object, actor string) (*core.LocalRelation, error) {
	lr := &core.LocalRelation{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   core.RelationTypeBlock,
	}
	queryString := "Actor = ? AND Object = ? AND Type = ?"
	queryParams := []any{actor, object, core.RelationTypeBlock}
	if id != "" {
		queryString = queryString + " AND ID = ?"
		queryParams = append(queryParams, id)
	}
	if tx := dao.Where(queryString, queryParams...).First(lr); tx.Error != nil {
		log.Printf("%s", tx.Error.Error())
		return lr, tx.Error
	}

	// mastodon
	if err := dao.Delete(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return lr, err
	}

	return lr, nil
}
