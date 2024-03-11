package handler

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func Like(o *orderedmap.OrderedMap) error {
	err := dao.Favourite(
		o.GetOrDefault("id", "").(string),
		o.GetOrDefault("object", "").(string),
		o.GetOrDefault("actor", "").(string),
	)
	return err
}

func Announce(o *orderedmap.OrderedMap) error {
	to := o.GetOrDefault("to", []string{}).([]string)
	cc := o.GetOrDefault("cc", []string{}).([]string)
	err := dao.Reblog(
		o.GetOrDefault("id", "").(string),
		o.GetOrDefault("object", "").(string),
		o.GetOrDefault("actor", "").(string),
		utils.ParseVisibility(to, cc),
	)
	return err
}

func UndoNotify(o *orderedmap.OrderedMap) error {
	switch o.GetOrDefault("type", "").(string) {
	case model.NotifyTypeLike:
		notify := &model.LocalNotify{
			Actor:  o.GetOrDefault("actor", "").(string),
			Object: o.GetOrDefault("object", "").(string),
			Type:   model.NotifyTypeLike,
		}
		if err := dao.Read(dao.DB(), notify); err != nil {
			return err
		}
		err := dao.Unreblog(notify)
		return err
	case model.NotifyTypeAnnounce:
		notify := &model.LocalNotify{
			Actor:  o.GetOrDefault("actor", "").(string),
			Object: o.GetOrDefault("object", "").(string),
			Type:   model.NotifyTypeAnnounce,
		}
		if err := dao.Read(dao.DB(), notify); err != nil {
			return err
		}
		err := dao.Unreblog(notify)
		return err
	default:
		log.Println("not support")
		return fmt.Errorf("not support")
	}
}
