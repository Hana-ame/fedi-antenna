package core

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Reblog(id, object, actor, visibility string) (*entities.Status, error) {
	notify := &core.LocalNotify{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   core.NotifyTypeAnnounce,

		Visibility: visibility,
	}

	reblog := convert.ToMastodonReblog(notify, false)

	if err := dao.Create(notify); err != nil {
		log.Printf("%s", err.Error())
		return reblog, err
	}

	// mastodon
	// status, err := c.ToMastodonReblog(notify, actor)
	return reblog, nil
}
func Unreblog(object, actor string) error {
	notify := &core.LocalNotify{
		Actor:  actor,
		Object: object,
		Type:   core.NotifyTypeAnnounce,
	}
	if err := dao.Read(notify); err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	if err := dao.Delete(notify); err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	return nil
}
