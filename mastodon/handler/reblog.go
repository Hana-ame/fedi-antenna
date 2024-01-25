package handler

import (
	"log"
	"strconv"

	c "github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Boost_a_status(id string, actor string, o *model.Boost_a_status) (*entities.Status, error) {
	status := &entities.Status{}
	if tx := dao.Where("Id = ?", id).First(status); tx.Error != nil {
		log.Printf("%s", tx.Error.Error())
		return nil, tx.Error
	}

	name, host := utils.ParseNameAndHost(actor)
	announceID := utils.ParseStatusesUri(name, host, strconv.Itoa(int(utils.Now()))) + "/activity"

	reblog, err := c.Reblog(announceID, status.Uri, actor, o.Visibility)
	if reblog != nil {
		reblog.Reblog = status
	}
	return reblog, err
}

func Undo_boost_of_a_status(id string, actor string) (*entities.Status, error) {
	status := &entities.Status{}
	if tx := dao.Where("Id = ?", id).First(status); tx.Error != nil {
		log.Printf("%s", tx.Error.Error())
		return nil, tx.Error
	}

	err := c.Unreblog(status.Uri, actor)

	// activitypub
	// todo

	return status, err
}
