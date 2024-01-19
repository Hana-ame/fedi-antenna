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
		//send again

		return convert.ToMastodonRelationship(id, actor), err
	}
	lr.ID = utils.GenerateObjectID(core.RelationTypeFollow, host)
	lr.Status = core.RelationStatusPadding

	// activitypub
	// todo
	// return if failed.

	// mastodon
	if err := dao.Create(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return nil, err
	}

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
	// return if failed.

	// mastodon
	if err := dao.Delete(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}
	if lr.Status == core.RelationStatusAccepted {
		dao.UpdateAccountFollowingCount(&entities.Account{Uri: lr.Actor}, -1)
		dao.UpdateAccountFollowersCount(&entities.Account{Uri: lr.Object}, -1)
	}

	return convert.ToMastodonRelationship(id, actor), nil
}

func View_pending_follow_requests(
	Authorization,
	max_id,
	since_id,
	limit string,
) (accts []*entities.Account, err error) {
	// todo
	// utils.Atoi(max_id, 0)
	tx := dao.DB().Limit(utils.Atoi(limit, 25)).Find(&accts)
	if tx.Error != nil {
		return accts, tx.Error
	}
	return accts, nil
}

func Accept_follow_request(
	id,
	actor string,
) (*entities.Relationship, error) {
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  id,
		Object: actor,
		Type:   core.RelationTypeFollow,
		Status: core.RelationStatusPadding,
	}
	if err := dao.Read(lr); err != nil {
		// 不存在的情况
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}

	// mastodon
	lr.Status = core.RelationStatusAccepted
	if err := dao.Update(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}
	if lr.Status == core.RelationStatusAccepted {
		dao.UpdateAccountFollowingCount(&entities.Account{Uri: lr.Actor}, 1)
		dao.UpdateAccountFollowersCount(&entities.Account{Uri: lr.Object}, 1)
	}

	// activitypub
	// todo
	// return if failed.

	return convert.ToMastodonRelationship(id, actor), nil
}
func Reject_follow_request(
	id,
	actor string,
) (*entities.Relationship, error) {
	lr := &core.LocalRelation{
		// ID:     utils.GenerateObjectID(typ, host),
		Actor:  id,
		Object: actor,
		Type:   core.RelationTypeFollow,
		Status: core.RelationStatusPadding,
	}
	if err := dao.Read(lr); err != nil {
		// 不存在的情况
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}

	// mastodon
	if err := dao.Delete(lr); err != nil {
		// should not
		log.Printf("%s", err.Error())
		return convert.ToMastodonRelationship(id, actor), err
	}

	// activitypub
	// todo
	// return if failed.

	return convert.ToMastodonRelationship(id, actor), nil
}
