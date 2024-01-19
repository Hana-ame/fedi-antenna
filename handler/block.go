package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Block_account(id, actor string) (*entities.Relationship, error) {
	{
		// delete follow id -> actor
		lr := &core.LocalRelation{
			// ID:     utils.GenerateObjectID(typ, host),
			Actor:  actor,
			Object: id,
			Type:   core.RelationTypeFollow,
			// Status: core.RelationStatusPadding,
		}
		if err := dao.Read(lr); err == nil {
			if err := dao.Delete(lr); err != nil {
				if lr.Status == core.RelationStatusAccepted {
					dao.UpdateAccountFollowingCount(&entities.Account{Uri: lr.Actor}, -1)
					dao.UpdateAccountFollowersCount(&entities.Account{Uri: lr.Object}, -1)
				}
			} else {
				return nil, err
			}
		}
	}
	{
		// delete follow actor -> id
		lr := &core.LocalRelation{
			// ID:     utils.GenerateObjectID(typ, host),
			Actor:  id,
			Object: actor,
			Type:   core.RelationTypeFollow,
			// Status: core.RelationStatusPadding,
		}
		if err := dao.Read(lr); err == nil {
			if err := dao.Delete(lr); err != nil {
				if lr.Status == core.RelationStatusAccepted {
					dao.UpdateAccountFollowingCount(&entities.Account{Uri: lr.Actor}, -1)
					dao.UpdateAccountFollowersCount(&entities.Account{Uri: lr.Object}, -1)
				}
			} else {
				return nil, err
			}
		}
	}
	_, host := utils.ParseNameAndHost(actor)
	lr := &core.LocalRelation{
		ID:     utils.GenerateObjectID(core.RelationTypeBlock, host),
		Actor:  actor,
		Object: id,
		Type:   core.RelationTypeBlock,
		Status: core.RelationStatusBlocked,
	}

	if err := dao.Read(lr); err == nil {
		// send again

		return convert.ToMastodonRelationship(id, actor), nil
	}

	if err := dao.Create(lr); err != nil {
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}

	return convert.ToMastodonRelationship(id, actor), nil

}

func Unblock_account(id, actor string) (*entities.Relationship, error) {
	lr := &core.LocalRelation{
		Actor:  actor,
		Object: id,
		Type:   core.RelationTypeBlock,
	}

	if err := dao.Read(lr); err != nil {
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}

	// activitypub
	// todo
	// return if failed.

	// mastodon
	if err := dao.Delete(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}

	return convert.ToMastodonRelationship(id, actor), nil
}
