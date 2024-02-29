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
		err := dao.Unfavourite(
			o.GetOrDefault("object", "").(string),
			o.GetOrDefault("actor", "").(string))
		return err
	case model.NotifyTypeAnnounce:
		err := dao.Unreblog(
			o.GetOrDefault("object", "").(string),
			o.GetOrDefault("actor", "").(string))
		return err
	default:
		log.Println("not support")
		return fmt.Errorf("not support")
	}
}
