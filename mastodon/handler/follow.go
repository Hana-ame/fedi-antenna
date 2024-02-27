package handler

import (
	"log"

	c "github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	accounts "github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Follow_account(id, actor string, o *accounts.Follow_account) (*entities.Relationship, error) {
	_, host := utils.ParseNameAndHost(actor)
	objectID := utils.GenerateObjectID(core.RelationTypeFollow, host)
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(acct); err != nil {
		log.Println(err)
		return convert.ToMastodonRelationship(id, actor), err
	}

	err := c.Follow(objectID, acct.Uri, actor)

	if err == nil {

	}

	return convert.ToMastodonRelationship(id, actor), err
}

func Unfollow_account(id, actor string) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(acct); err != nil {
		log.Println(err)
		return convert.ToMastodonRelationship(id, actor), err
	}

	err := c.Unfollow("", acct.Uri, actor)

	if err == nil {

	}

	return convert.ToMastodonRelationship(id, actor), err
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
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(acct); err != nil {
		log.Println(err)
		return convert.ToMastodonRelationship(id, actor), err
	}
	err := c.Accept("", acct.Uri, actor)

	if err == nil {

	}

	return convert.ToMastodonRelationship(id, actor), err
}
func Reject_follow_request(
	id,
	actor string,
) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(acct); err != nil {
		log.Println(err)
		return convert.ToMastodonRelationship(id, actor), err
	}

	err := c.Reject("", acct.Uri, actor)

	if err == nil {

	}

	return convert.ToMastodonRelationship(id, actor), err
}
