package core

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Follow(id, object, actor string) error {
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  actor,
		Object: object,
		Type:   core.RelationTypeFollow,
		// Status: core.RelationStatusPadding,
	}
	if err := dao.Read(lr); err == nil {
		if lr.Status == core.RelationStatusPadding {
			//send again
		}
		return err
	}
	lr.Status = core.RelationStatusPadding
	lr.ID = id
	// activitypub
	// todo
	// return if failed.

	// mastodon
	if err := dao.Create(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return err
	}
	return nil
}

func Unfollow(id, object, actor string) error {
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  actor,
		Object: object,
		Type:   core.RelationTypeFollow,
		// Status: core.RelationStatusPadding,
	}
	if err := dao.Read(lr); err != nil {
		// 不存在的情况
		log.Printf("%s", err.Error())
		return err
	}

	// activitypub
	// todo
	// return if failed.

	// mastodon
	if err := dao.Delete(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return err
	}
	if lr.Status == core.RelationStatusAccepted {
		dao.UpdateAccountFollowingCount(&entities.Account{Uri: lr.Actor}, -1)
		dao.UpdateAccountFollowersCount(&entities.Account{Uri: lr.Object}, -1)
	}

	return nil
}

func Accept(id, object, actor string) error {
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  object,
		Object: actor,
		Type:   core.RelationTypeFollow,
		Status: core.RelationStatusPadding,
	}
	if err := dao.Read(lr); err != nil {
		// 不存在的情况
		log.Printf("%s", err.Error())
		return err
	}

	// mastodon
	lr.Status = core.RelationStatusAccepted
	if err := dao.Update(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return err
	}
	if lr.Status == core.RelationStatusAccepted {
		dao.UpdateAccountFollowingCount(&entities.Account{Uri: lr.Actor}, 1)
		dao.UpdateAccountFollowersCount(&entities.Account{Uri: lr.Object}, 1)
	}

	// activitypub
	// todo
	// return if failed.

	return nil
}

func Reject(id, object, actor string) error {
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  object,
		Object: actor,
		Type:   core.RelationTypeFollow,
		Status: core.RelationStatusPadding,
	}
	if err := dao.Read(lr); err != nil {
		// 不存在的情况
		log.Printf("%s", err.Error())
		return err
	}

	// mastodon
	if err := dao.Delete(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return err
	}

	// activitypub
	// todo
	// return if failed.
	return nil
}
