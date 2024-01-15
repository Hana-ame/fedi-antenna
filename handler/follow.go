package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	accounts "github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Follow_account(id, actor string, o *accounts.Follow_account) (*entities.Relationship, error) {
	_, host := utils.ParseNameAndHost(actor)
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  actor,
		Object: id,
		Type:   core.RelationTypeFollow,
		// Status: core.RelationStatusPadding,
	}
	if err := dao.Read(lr); err == nil {
		return convert.ToMastodonRelationship(id, actor), err
	}

	lr.ID = utils.GenerateObjectID(core.RelationTypeFollow, host)
	lr.Status = core.RelationStatusPadding
	if err := dao.Create(lr); err != nil {
		return nil, err
	}

	// activitypub
	// todo

	// mastodon
	return convert.ToMastodonRelationship(id, actor), nil
}

func Unfollow_account(id, actor string) (*entities.Relationship, error) {
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  actor,
		Object: id,
		Type:   core.RelationTypeFollow,
		// Status: core.RelationStatusPadding,
	}
	if err := dao.Read(lr); err != nil {
		// 不存在的情况
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}
	// activitypub
	// todo

	// mastodon
	if err := dao.Delete(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}

	return convert.ToMastodonRelationship(id, actor), nil
}
