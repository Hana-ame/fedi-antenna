package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	accounts "github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Follow_account(id, actor string, o *accounts.Follow_account) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(dao.DB(), acct); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	_, host := utils.ActivitypubID2NameAndHost(actor)
	objectID := utils.NewObjectID(model.RelationTypeFollow, host)
	if err := dao.Follow(objectID, acct.Uri, actor); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	_, host = utils.ActivitypubID2NameAndHost(acct.Uri)
	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	return dao.Relationship(acct.Uri, actor)
}

func Unfollow_account(id, actor string) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(dao.DB(), acct); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	_, host := utils.ActivitypubID2NameAndHost(actor)
	if err := dao.Unfollow("", acct.Uri, actor); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	_, host = utils.ActivitypubID2NameAndHost(acct.Uri)
	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	return dao.Relationship(acct.Uri, actor)
}

// TODO
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

func Accept_follow_request(id, actor string) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(dao.DB(), acct); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	if err := dao.Accept("", acct.Uri, actor); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	_, host := utils.ActivitypubID2NameAndHost(acct.Uri)
	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	return dao.Relationship(acct.Uri, actor)
}

func Reject_follow_request(id, actor string) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(dao.DB(), acct); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	if err := dao.Reject("", acct.Uri, actor); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	_, host := utils.ActivitypubID2NameAndHost(acct.Uri)
	if core.IsLocal(host) {
		// if host not at local
		// then post to remote. TODO
	}

	return dao.Relationship(acct.Uri, actor)
}
