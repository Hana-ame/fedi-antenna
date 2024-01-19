package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	accounts "github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Block_account(id, actor string, o *accounts.Follow_account) (*entities.Relationship, error) {
	// _, host := utils.ParseNameAndHost(actor)
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  actor,
		Object: id,
		Type:   core.RelationTypeBlock,
		// Status: core.RelationStatusPadding,
	}

	if err := dao.Read(lr); err == nil {
		// 再次发送

		return convert.ToMastodonRelationship(id, actor), err
	} else {
		lr.Status = core.RelationStatusBlocked

		if err := dao.Create(lr); err != nil {
			log.Printf("%s", err.Error())
			return nil, err
		} else {
			if err := dao.Delete(&core.LocalRelation{
				Actor:  actor,
				Object: id,
				Type:   core.RelationTypeFollow,
			}); err != nil {
				// dao.
			}
			dao.Delete(&core.LocalRelation{
				Actor:  id,
				Object: actor,
				Type:   core.RelationTypeFollow,
			})
		}

		return convert.ToMastodonRelationship(id, actor), nil
	}
}

func Unblock_account(id, actor string) (*entities.Relationship, error) {
	lr := &core.LocalRelation{
		Actor:  actor,
		Object: id,
		Type:   core.RelationTypeBlock,
	}
	// 不存在的情况
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
