package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/actions"
	c "github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Block_account(id, actor string) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(acct); err != nil {
		log.Println(err)
		return convert.ToMastodonRelationship(id, actor), err
	}

	_, host := utils.ParseNameAndHost(actor)
	err := c.Block(utils.GenerateObjectID(model.RelationTypeBlock, host), acct.Uri, actor)

	if err == nil {
		go actions.Block(actor, acct.Uri)
	}

	return convert.ToMastodonRelationship(acct.Uri, actor), err
}

func Unblock_account(id, actor string) (*entities.Relationship, error) {
	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(acct); err != nil {
		log.Println(err)
		return convert.ToMastodonRelationship(id, actor), err
	}

	relation, err := c.Unblock("", acct.Uri, actor)

	if err == nil {
		go actions.UndoBlock(actor, relation.ID)
	}

	return convert.ToMastodonRelationship(acct.Uri, actor), err
}
