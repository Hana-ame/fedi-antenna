package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/actions"
	c "github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// id is the mastodon published which is timestamp in us
func Favourite_a_status(id string, actor string) (*entities.Status, error) {
	status := &entities.Status{
		Id: id,
	}
	if tx := dao.Where("Id = ?", id).First(status); tx.Error != nil {
		log.Printf("%s", tx.Error.Error())
		return nil, tx.Error
	}

	_, host := utils.ParseNameAndHost(actor)
	err := c.Favourite(utils.GenerateObjectID("favourite", host), status.Uri, actor)

	if err == nil {
		go actions.Like(actor, status.Uri)
	}

	return status, err
}

func Undo_favourite_of_a_status(id string, actor string) (*entities.Status, error) {
	status := &entities.Status{
		Id: id,
	}
	if tx := dao.Where("Id = ?", id).First(status); tx.Error != nil {
		log.Printf("%s", tx.Error.Error())
		return nil, tx.Error
	}
	// mastodon
	err := c.Unfavourite(status.Uri, actor)

	if err == nil {
		// go actions.Like(actor, status.Uri)
	}

	return status, err
}
