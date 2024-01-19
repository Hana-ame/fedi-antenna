package handler

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	c "github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Block_account(id, actor string) (*entities.Relationship, error) {
	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(status); err != nil {
		return nil, err
	}

	_, host := utils.ParseNameAndHost(actor)
	err := c.Block(utils.GenerateObjectID(model.RelationTypeBlock, host), status.Uri, actor)

	if err == nil {
		go actions.Block(actor, status.Uri)
	}

	return convert.ToMastodonRelationship(status.Uri, actor), err
}

func Unblock_account(id, actor string) (*entities.Relationship, error) {
	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(status); err != nil {
		return nil, err
	}

	relation, err := c.Unblock("", status.Uri, actor)

	if err == nil {
		go actions.UndoBlock(actor, relation.ID)
	}

	return convert.ToMastodonRelationship(status.Uri, actor), err
}
