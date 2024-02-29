package handler

import (
	"log"
	"strconv"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Boost_a_status(id string, actor string, o *mastodon.Boost_a_status) (*entities.Status, error) {
	tx := dao.DB()

	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	name, host := utils.ActivitypubID2NameAndHost(actor)
	announceID := utils.ParseStatusesUri(name, host, strconv.Itoa(int(utils.NewTimestamp()))) + "/activity"
	if err := dao.Reblog(announceID, status.Uri, actor, o.Visibility); err != nil {
		log.Println(err)
		return nil, err
	}

	_, host = utils.ActivitypubID2NameAndHost(status.Uri)
	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	localNotify := &model.LocalNotify{
		ID: announceID,
	}
	if err := dao.Read(tx, localNotify); err != nil {
		return nil, err
	}

	return convert.ToMastodonReblog(localNotify, status), nil
}

func Undo_boost_of_a_status(id string, actor string) (*entities.Status, error) {
	tx := dao.DB()

	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	if err := dao.Unreblog(status.Uri, actor); err != nil {
		log.Println(err)
		return status, err
	}

	_, host := utils.ActivitypubID2NameAndHost(status.Uri)
	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	return status, nil
}
