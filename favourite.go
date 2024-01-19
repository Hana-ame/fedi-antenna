package core

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
)

func Favourite(id, object, actor string) error {
	favourite := &core.LocalNotify{
		ID:     id,
		Actor:  actor,
		Object: object,
		Type:   core.NotifyTypeLike,
	}

	if err := dao.Create(favourite); err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	// mastodon
	return nil
}

func Unfavourite(object, actor string) error {
	favourite := &core.LocalNotify{
		Actor:  actor,
		Object: object,
		Type:   core.NotifyTypeLike,
	}
	if err := dao.Read(favourite); err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	if err := dao.Delete(favourite); err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	// mastodon
	return nil
}
