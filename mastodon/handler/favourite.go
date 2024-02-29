package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// id is the mastodon published which is timestamp in us
func Favourite_a_status(id string, actor string) (*entities.Status, error) {
	tx := dao.Begin()

	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	_, host := utils.ActivitypubID2NameAndHost(actor)
	if err := dao.Favourite(utils.NewObjectID("favourite", host), status.Uri, actor); err != nil {
		log.Println(err)
		return status, err
	}

	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	return status, nil
}

func Undo_favourite_of_a_status(id string, actor string) (*entities.Status, error) {
	tx := dao.Begin()

	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	_, host := utils.ActivitypubID2NameAndHost(actor)
	if err := dao.Unfavourite(status.Uri, actor); err != nil {
		log.Println(err)
		return status, err
	}

	core.IsLocal(host)
	// if host not at local
	// then post to remote. TODO

	return status, nil
}
