package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// id is a number in string
func Block_account(id, actor string) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(dao.DB(), acct); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	_, host := utils.ActivitypubID2NameAndHost(actor)
	if err := dao.Block(utils.NewObjectID(model.RelationTypeBlock, host), acct.Uri, actor); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	return dao.Relationship(acct.Uri, actor)
}

func Unblock_account(id, actor string) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(dao.DB(), acct); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	_, host := utils.ActivitypubID2NameAndHost(actor)
	if err := dao.Unblock("", acct.Uri, actor); err != nil {
		log.Println(err)
		relationship, _ := dao.Relationship(acct.Uri, actor)
		return relationship, err
	}

	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	return dao.Relationship(acct.Uri, actor)
}
