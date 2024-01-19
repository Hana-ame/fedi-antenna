package handler

import (
	"log"
	"strconv"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Boost_a_status(id string, actor string, o *model.Boost_a_status) (*entities.Status, error) {
	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(status); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	name, host := utils.ParseNameAndHost(actor)
	announceID := utils.ParseStatusesUri(name, host, strconv.Itoa(int(utils.Now()))) + "/activity"

	notify := &core.LocalNotify{
		ID:     announceID,
		Actor:  actor,
		Object: status.Uri,
		Type:   core.NotifyTypeAnnounce,

		Visibility: o.Visibility,
	}
	if err := dao.Create(notify); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	// activitypub
	// todo

	// mastodon
	// status, err := c.ToMastodonReblog(notify, actor)

	// return status, err
	return nil, nil
}

func Undo_boost_of_a_status(id string, actor string) (*entities.Status, error) {
	published, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	ln := &core.LocalNote{
		Published: int64(published),
	}
	if err := dao.Read(ln); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	notify := &core.LocalNotify{
		Actor:  actor,
		Object: ln.ID,
		Type:   core.NotifyTypeAnnounce,
	}
	if err := dao.Read(notify); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	if err := dao.Delete(ln); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	// activitypub
	// todo

	// mastodon
	status := convert.ToMastodonReblog(notify)

	return status, nil
}
